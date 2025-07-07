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

func (s *PhoneService) GetPhone(userID primitive.ObjectID, phoneID primitive.ObjectID) (*model.PhoneNumber, error) {
	ctx := context.Background()
	phone, err := s.phoneRepo.FindPhoneByID(ctx, phoneID)
	if err != nil {
		return nil, err
	}
	// Verify the phone belongs to the user
	if phone.UserID != userID {
		return nil, nil // Or return a proper error for unauthorized access
	}
	return phone, nil
}

func (s *PhoneService) UpdatePhone(userID primitive.ObjectID, phoneID primitive.ObjectID, updatedPhone *model.PhoneNumber) error {
	ctx := context.Background()
	// First verify the phone belongs to the user
	phone, err := s.phoneRepo.FindPhoneByID(ctx, phoneID)
	if err != nil {
		return err
	}
	if phone.UserID != userID {
		return nil // Or return a proper error for unauthorized access
	}

	// Update the phone with new data
	updatedPhone.ID = phoneID
	updatedPhone.UserID = userID

	if err := s.phoneRepo.UpdatePhone(ctx, updatedPhone); err != nil {
		return err
	}

	return nil
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

	return nil
}
