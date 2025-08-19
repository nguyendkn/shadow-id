package repositories

import (
	"context"

	"shadow-id/internal/domain/entities"
	"shadow-id/pkg/types"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entities.User) error
	
	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id types.ID) (*entities.User, error)
	
	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	
	// Update updates an existing user
	Update(ctx context.Context, user *entities.User) error
	
	// Delete deletes a user by ID
	Delete(ctx context.Context, id types.ID) error
	
	// List retrieves all users with pagination
	List(ctx context.Context, limit, offset int) ([]*entities.User, error)
	
	// Count returns the total number of users
	Count(ctx context.Context) (int64, error)
}
