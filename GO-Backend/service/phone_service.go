package service

import (
	"context"
	model "go-fiber-app/models"
	"go-fiber-app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhoneService struct {
	phoneRepo *repository.PhoneRepository
	userRepo  *repository.UserRepository
}

func NewPhoneService(phoneRepo *repository.PhoneRepository, userRepo *repository.UserRepository) *PhoneService {
	return &PhoneService{phoneRepo: phoneRepo, userRepo: userRepo}
}

func (s *PhoneService) CreatePhone(phone *model.PhoneNumber) error {
	if !phone.Validate() {
		return nil // Add proper validation error handling if needed
	}
	ctx := context.Background()

	// Validate that the UserID exists
	if _, err := s.userRepo.FindUserByID(ctx, phone.UserID); err != nil {
		return err // Return error if user not found
	}

	// Create the phone number
	if err := s.phoneRepo.CreatePhone(ctx, phone); err != nil {
		return err
	}

	// Update the user's PhoneIDs with the new phone ID
	user, err := s.userRepo.FindUserByID(ctx, phone.UserID)
	if err != nil {
		return err
	}
	user.PhoneIDs = append(user.PhoneIDs, phone.ID)
	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *PhoneService) GetPhonesByUser(userID primitive.ObjectID) ([]model.PhoneNumber, error) {
	ctx := context.Background()
	phones, err := s.phoneRepo.GetPhonesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return phones, nil
}

func (s *PhoneService) DeletePhone(userID primitive.ObjectID, phoneID primitive.ObjectID) error {
	ctx := context.Background()

	// First verify the phone belongs to the user
	phone, err := s.phoneRepo.FindPhoneByID(ctx, phoneID)
	if err != nil {
		return err
	}
	if phone.UserID != userID {
		return nil // Or return a proper error for unauthorized access
	}

	// Delete the phone
	if err := s.phoneRepo.DeletePhone(ctx, phoneID); err != nil {
		return err
	}

	// Remove phone ID from user's PhoneIDs array
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Remove phoneID from user's PhoneIDs slice
	for i, id := range user.PhoneIDs {
		if id == phoneID {
			user.PhoneIDs = append(user.PhoneIDs[:i], user.PhoneIDs[i+1:]...)
			break
		}
	}

	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *PhoneService) GetUserWithPhones(userID primitive.ObjectID) (*model.User, error) {
	ctx := context.Background()
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
