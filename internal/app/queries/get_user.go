package queries

import (
	"context"

	"shadow-id/internal/domain/repositories"
	"shadow-id/pkg/errors"
	"shadow-id/pkg/types"
)

// GetUserQuery represents the query to get a user
type GetUserQuery struct {
	ID types.ID `json:"id" validate:"required"`
}

// GetUserResult represents the result of getting a user
type GetUserResult struct {
	ID        types.ID `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// GetUserHandler handles the get user query
type GetUserHandler struct {
	userRepo repositories.UserRepository
}

// NewGetUserHandler creates a new get user handler
func NewGetUserHandler(userRepo repositories.UserRepository) *GetUserHandler {
	return &GetUserHandler{
		userRepo: userRepo,
	}
}

// Handle executes the get user query
func (h *GetUserHandler) Handle(ctx context.Context, query GetUserQuery) (*GetUserResult, error) {
	// Get user from repository
	user, err := h.userRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	
	if user == nil {
		return nil, errors.NewNotFoundError("user not found")
	}
	
	// Return result
	return &GetUserResult{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
