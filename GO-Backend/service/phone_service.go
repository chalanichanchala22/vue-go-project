package service

import (
	"context"
	"fmt"
	model "go-fiber-app/models"
	"go-fiber-app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhoneService struct {
	phoneRepo *repository.PhoneRepository
}

func NewPhoneService(phoneRepo *repository.PhoneRepository) *PhoneService {
	return &PhoneService{phoneRepo: phoneRepo}
}

func (s *PhoneService) CreatePhone(phone *model.PhoneNumber) error {
	fmt.Printf("PhoneService.CreatePhone called with: %+v\n", phone)

	if !phone.Validate() {
		fmt.Printf("Phone validation failed for: %+v\n", phone)
		return fmt.Errorf("phone validation failed")
	}

	fmt.Printf("Phone validation passed, calling repository...\n")
	ctx := context.Background()
	if err := s.phoneRepo.CreatePhone(ctx, phone); err != nil {
		fmt.Printf("Repository error: %v\n", err)
		return err
	}

	fmt.Printf("Phone successfully created in repository\n")
	return nil
}

func (s *PhoneService) GetPhonesByUser(userID primitive.ObjectID) ([]*model.PhoneNumber, error) {
	ctx := context.Background()
	return s.phoneRepo.GetPhonesByUser(ctx, userID)
}

func (s *PhoneService) UpdatePhone(phone *model.PhoneNumber) error {
	if !phone.Validate() {
		return fmt.Errorf("phone validation failed")
	}
	ctx := context.Background()
	if err := s.phoneRepo.UpdatePhone(ctx, phone); err != nil {
		return err
	}
	return nil
}

func (s *PhoneService) DeletePhone(userID, phoneID primitive.ObjectID) error {
	ctx := context.Background()
	if err := s.phoneRepo.DeletePhone(ctx, userID, phoneID); err != nil {
		return err
	}
	return nil
}
