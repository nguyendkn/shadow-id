package services

import (
	"context"

	"shadow-id/internal/domain/entities"
	"shadow-id/pkg/types"
)

// UserService defines domain services for user operations
type UserService interface {
	// ValidateUserCreation validates user creation business rules
	ValidateUserCreation(ctx context.Context, user *entities.User) error
	
	// ValidateUserUpdate validates user update business rules
	ValidateUserUpdate(ctx context.Context, user *entities.User) error
	
	// IsEmailUnique checks if email is unique
	IsEmailUnique(ctx context.Context, email string, excludeID types.ID) (bool, error)
}
