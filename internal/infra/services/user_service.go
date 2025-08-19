package services

import (
	"context"
	"strings"

	"shadow-id/internal/domain/entities"
	"shadow-id/internal/domain/repositories"
	"shadow-id/pkg/types"
)

// UserService implements domain user service
type UserService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// ValidateUserCreation validates user creation business rules
func (s *UserService) ValidateUserCreation(ctx context.Context, user *entities.User) error {
	// Business rule: Name should not contain only whitespace
	if strings.TrimSpace(user.Name) == "" {
		return entities.ErrInvalidUserName
	}
	
	// Business rule: Email should be properly formatted (basic check)
	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return entities.ErrInvalidUserEmail
	}
	
	return nil
}

// ValidateUserUpdate validates user update business rules
func (s *UserService) ValidateUserUpdate(ctx context.Context, user *entities.User) error {
	// Same validation as creation for now
	return s.ValidateUserCreation(ctx, user)
}

// IsEmailUnique checks if email is unique
func (s *UserService) IsEmailUnique(ctx context.Context, email string, excludeID types.ID) (bool, error) {
	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	
	// Email is unique if no user found
	if existingUser == nil {
		return true, nil
	}
	
	// Email is unique if the existing user is the one being excluded (for updates)
	if excludeID != "" && existingUser.ID == excludeID {
		return true, nil
	}
	
	return false, nil
}
