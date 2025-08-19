package wails

import (
	"context"

	"shadow-id/internal/app/commands"
	"shadow-id/internal/app/queries"
	"shadow-id/internal/app/services"
	"shadow-id/internal/infra/config"
	infraservices "shadow-id/internal/infra/services"
	"shadow-id/internal/infra/storage/memory"
	"shadow-id/pkg/logger"
	"shadow-id/pkg/types"
)

// App struct represents the Wails application
type App struct {
	ctx    context.Context
	config *config.Config
	logger logger.Logger

	// Application services
	appService *services.ApplicationService
}

// NewApp creates a new Wails application instance
func NewApp() (*App, error) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Initialize logger
	appLogger := logger.New(cfg.LogLevel)

	// Initialize repositories
	userRepo := memory.NewUserRepository()

	// Initialize domain services
	userService := infraservices.NewUserService(userRepo)

	// Initialize application services
	appService := services.NewApplicationService(userRepo, userService)

	return &App{
		config:     cfg,
		logger:     appLogger,
		appService: appService,
	}, nil
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Info("Application started successfully")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	a.logger.Info("Greet method called", "name", name)
	return "Hello " + name + ", It's show time!"
}

// CreateUser creates a new user
func (a *App) CreateUser(name, email string) (*commands.CreateUserResult, error) {
	a.logger.Info("CreateUser method called", "name", name, "email", email)

	cmd := commands.CreateUserCommand{
		Name:  name,
		Email: email,
	}

	result, err := a.appService.Commands.CreateUser.Handle(a.ctx, cmd)
	if err != nil {
		a.logger.Error("Failed to create user", "error", err)
		return nil, err
	}

	a.logger.Info("User created successfully", "id", result.ID)
	return result, nil
}

// GetUser retrieves a user by ID
func (a *App) GetUser(id string) (*queries.GetUserResult, error) {
	a.logger.Info("GetUser method called", "id", id)

	userID := types.ID(id)
	query := queries.GetUserQuery{
		ID: userID,
	}

	result, err := a.appService.Queries.GetUser.Handle(a.ctx, query)
	if err != nil {
		a.logger.Error("Failed to get user", "error", err)
		return nil, err
	}

	a.logger.Info("User retrieved successfully", "id", result.ID)
	return result, nil
}

// GetAppInfo returns application information
func (a *App) GetAppInfo() map[string]interface{} {
	return map[string]interface{}{
		"name":    a.config.AppName,
		"version": a.config.Version,
		"env":     a.config.Environment,
	}
}
