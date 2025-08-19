package services

import (
	"shadow-id/internal/app/commands"
	"shadow-id/internal/app/queries"
	"shadow-id/internal/domain/repositories"
	"shadow-id/internal/domain/services"
)

// ApplicationService aggregates all application services
type ApplicationService struct {
	Commands *CommandHandlers
	Queries  *QueryHandlers
}

// CommandHandlers aggregates all command handlers
type CommandHandlers struct {
	CreateUser *commands.CreateUserHandler
}

// QueryHandlers aggregates all query handlers
type QueryHandlers struct {
	GetUser *queries.GetUserHandler
}

// NewApplicationService creates a new application service
func NewApplicationService(
	userRepo repositories.UserRepository,
	userService services.UserService,
) *ApplicationService {
	return &ApplicationService{
		Commands: &CommandHandlers{
			CreateUser: commands.NewCreateUserHandler(userRepo, userService),
		},
		Queries: &QueryHandlers{
			GetUser: queries.NewGetUserHandler(userRepo),
		},
	}
}
