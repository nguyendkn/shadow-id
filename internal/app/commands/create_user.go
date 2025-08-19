package commands

import (
	"context"

	"shadow-id/internal/domain/entities"
	"shadow-id/internal/domain/repositories"
	"shadow-id/internal/domain/services"
	"shadow-id/pkg/errors"
	"shadow-id/pkg/types"
)

// CreateUserCommand represents the command to create a user
type CreateUserCommand struct {
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
}

// CreateUserResult represents the result of creating a user
type CreateUserResult struct {
	ID        types.ID `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	CreatedAt string   `json:"created_at"`
}

// CreateUserHandler handles the create user command
type CreateUserHandler struct {
	userRepo    repositories.UserRepository
	userService services.UserService
}

// NewCreateUserHandler creates a new create user handler
func NewCreateUserHandler(
	userRepo repositories.UserRepository,
	userService services.UserService,
) *CreateUserHandler {
	return &CreateUserHandler{
		userRepo:    userRepo,
		userService: userService,
	}
}

// Handle executes the create user command
func (h *CreateUserHandler) Handle(ctx context.Context, cmd CreateUserCommand) (*CreateUserResult, error) {
	// Create user entity
	user := entities.NewUser(cmd.Name, cmd.Email)
	
	// Validate user
	if err := user.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid user data")
	}
	
	// Validate business rules
	if err := h.userService.ValidateUserCreation(ctx, user); err != nil {
		return nil, errors.Wrap(err, "user creation validation failed")
	}
	
	// Check email uniqueness
	isUnique, err := h.userService.IsEmailUnique(ctx, user.Email, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to check email uniqueness")
	}
	if !isUnique {
		return nil, entities.ErrUserAlreadyExists
	}
	
	// Save user
	if err := h.userRepo.Create(ctx, user); err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	
	// Return result
	return &CreateUserResult{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
