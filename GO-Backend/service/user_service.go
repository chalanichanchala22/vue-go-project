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

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) SetPhoneRepository(phoneRepo *repository.PhoneRepository) {
	s.phoneRepo = phoneRepo
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

	// Get the user's phone numbers if phone repository is available
	if s.phoneRepo != nil {
		phones, err := s.phoneRepo.GetPhonesByUser(ctx, id)
		if err != nil {
			// Log the error but don't fail the request
			fmt.Printf("Error fetching phones for user %s: %v\n", id.Hex(), err)
			// Initialize empty slice even on error
			user.Phones = []*model.PhoneNumber{}
		} else {
			user.Phones = phones
		}
	} else {
		// Initialize empty phones slice if no phone repository
		user.Phones = []*model.PhoneNumber{}
	}

	return user, nil
}

func (s *UserService) GetAllUsersWithPhones() ([]*model.User, error) {
	ctx := context.Background()
	// Get all users
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Get phone numbers for each user if phone repository is available
	if s.phoneRepo != nil {
		for _, user := range users {
			phones, err := s.phoneRepo.GetPhonesByUser(ctx, user.ID)
			if err != nil {
				// Log the error but don't fail the request
				fmt.Printf("Error fetching phones for user %s: %v\n", user.ID.Hex(), err)
				// Initialize empty slice even on error
				user.Phones = []*model.PhoneNumber{}
			} else {
				user.Phones = phones
			}
		}
	} else {
		// Initialize empty phones slice for all users if no phone repository
		for _, user := range users {
			user.Phones = []*model.PhoneNumber{}
		}
	}

	return users, nil
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
	return nil
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	ctx := context.Background()
	return s.userRepo.FindUserByEmail(ctx, email)
}
