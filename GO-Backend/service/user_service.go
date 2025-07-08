package service

import (
	"context"
	"fmt"
	model "go-fiber-app/models"
	"go-fiber-app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo  *repository.UserRepository
	phoneRepo *repository.PhoneRepository
}

func NewUserService(userRepo *repository.UserRepository, phoneRepo *repository.PhoneRepository) *UserService {
	return &UserService{userRepo: userRepo, phoneRepo: phoneRepo}
}

func (s *UserService) CreateUser(user *model.User) error {
	if !user.Validate() {
		return fmt.Errorf("user validation failed")
	}
	ctx := context.Background()
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetAllUsers() ([]*model.User, error) {
	ctx := context.Background()
	return s.userRepo.GetAllUsers(ctx)
}

func (s *UserService) GetUser(id primitive.ObjectID) (*model.User, error) {
	ctx := context.Background()
	return s.userRepo.FindUserByID(ctx, id)
}

func (s *UserService) GetUserWithPhones(id primitive.ObjectID) (*model.User, error) {
	ctx := context.Background()

	// Get the user
	user, err := s.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Get the user's phones
	phones, err := s.phoneRepo.GetPhonesByUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// Add phones to the user
	user.Phones = phones

	return user, nil
}

func (s *UserService) UpdateUser(user *model.User) error {
	ctx := context.Background()
	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id primitive.ObjectID) error {
	ctx := context.Background()
	if err := s.userRepo.DeleteUser(ctx, id); err != nil {
		return err
	}
	return s.phoneRepo.DeletePhonesByUserID(ctx, id)
}
