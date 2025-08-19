# Shadow ID - Clean Architecture Documentation

## Overview

Shadow ID is built using Clean Architecture principles, ensuring separation of concerns, testability, and maintainability. The application follows the dependency inversion principle where high-level modules don't depend on low-level modules.

## Architecture Layers

### 1. Domain Layer (`internal/domain/`)

The core business logic layer that contains:

- **Entities** (`entities/`): Core business objects (User, etc.)
- **Repositories** (`repositories/`): Interfaces for data access
- **Services** (`services/`): Domain services for business rules

**Dependencies**: None (pure business logic)

### 2. Application Layer (`internal/app/`)

The use case layer that orchestrates business logic:

- **Commands** (`commands/`): Write operations (Create, Update, Delete)
- **Queries** (`queries/`): Read operations (Get, List)
- **Services** (`services/`): Application services that coordinate use cases

**Dependencies**: Domain layer only

### 3. Infrastructure Layer (`internal/infra/`)

The implementation layer that provides concrete implementations:

- **Wails** (`wails/`): Wails application integration
- **Storage** (`storage/`): Repository implementations
- **Services** (`services/`): Domain service implementations
- **Config** (`config/`): Configuration management

**Dependencies**: Application and Domain layers

## Directory Structure

```
shadow-id/
├── cmd/                    # Application entry points
│   └── shadow-id/
│       └── main.go
├── internal/               # Private application code
│   ├── app/               # Application layer
│   │   ├── commands/      # Command handlers (CQRS)
│   │   ├── queries/       # Query handlers (CQRS)
│   │   └── services/      # Application services
│   ├── domain/            # Domain layer
│   │   ├── entities/      # Business entities
│   │   ├── repositories/  # Repository interfaces
│   │   └── services/      # Domain service interfaces
│   └── infra/             # Infrastructure layer
│       ├── wails/         # Wails integration
│       ├── storage/       # Data storage implementations
│       ├── services/      # Domain service implementations
│       └── config/        # Configuration
├── pkg/                   # Public packages
│   ├── errors/           # Error handling utilities
│   ├── logger/           # Logging utilities
│   └── types/            # Common types
├── configs/              # Configuration files
├── frontend/             # Next.js frontend
└── build/               # Build artifacts
```

## Design Patterns

### 1. CQRS (Command Query Responsibility Segregation)

- **Commands**: Handle write operations (Create, Update, Delete)
- **Queries**: Handle read operations (Get, List)
- Separate models for read and write operations

### 2. Repository Pattern

- Abstract data access through interfaces
- Concrete implementations in infrastructure layer
- Easy to swap storage backends

### 3. Dependency Injection

- Dependencies injected through constructors
- Interfaces used for loose coupling
- Easy testing with mocks

### 4. Error Handling

- Custom error types with context
- Wrapped errors for better debugging
- Consistent error responses

## Key Principles

### 1. Dependency Inversion

```
Infrastructure → Application → Domain
```

- Domain layer has no dependencies
- Application layer depends only on Domain
- Infrastructure layer depends on Application and Domain

### 2. Single Responsibility

Each component has a single, well-defined responsibility:

- Entities: Business data and rules
- Repositories: Data access
- Services: Business logic coordination
- Handlers: Use case execution

### 3. Interface Segregation

Small, focused interfaces:

- Repository interfaces per entity
- Service interfaces per domain
- Handler interfaces per use case

## Testing Strategy

### Unit Tests

- Domain entities and services
- Application handlers
- Infrastructure services

### Integration Tests

- Repository implementations
- Wails application integration
- End-to-end workflows

### Test Structure

```
internal/
├── domain/
│   ├── entities/
│   │   ├── user.go
│   │   └── user_test.go
│   └── services/
│       ├── user_service.go
│       └── user_service_test.go
└── app/
    ├── commands/
    │   ├── create_user.go
    │   └── create_user_test.go
    └── queries/
        ├── get_user.go
        └── get_user_test.go
```

## Configuration Management

### Environment Variables

- `APP_NAME`: Application name
- `APP_VERSION`: Application version
- `APP_ENV`: Environment (development, production, test)
- `LOG_LEVEL`: Logging level (debug, info, warn, error)

### Configuration Files

- `configs/app.yaml`: Default configuration
- Environment-specific overrides supported

## Error Handling

### Error Types

- `ValidationError`: Input validation errors
- `NotFoundError`: Resource not found
- `ConflictError`: Business rule violations
- `InternalError`: System errors

### Error Wrapping

```go
if err := repo.Create(ctx, user); err != nil {
    return errors.Wrap(err, "failed to create user")
}
```

## Logging

### Structured Logging

```go
logger.Info("User created", "id", user.ID, "email", user.Email)
logger.Error("Failed to create user", "error", err)
```

### Log Levels

- `DEBUG`: Detailed debugging information
- `INFO`: General information
- `WARN`: Warning messages
- `ERROR`: Error messages

## Future Enhancements

### Database Integration

- PostgreSQL repository implementation
- Database migrations
- Connection pooling

### API Layer

- REST API endpoints
- GraphQL support
- API documentation

### Authentication & Authorization

- JWT token handling
- Role-based access control
- OAuth integration

### Monitoring & Observability

- Metrics collection
- Distributed tracing
- Health checks

## Getting Started

1. **Install Dependencies**
   ```bash
   go mod tidy
   ```

2. **Run Application**
   ```bash
   wails dev
   ```

3. **Run Tests**
   ```bash
   go test ./...
   ```

4. **Build Application**
   ```bash
   wails build
   ```

## Contributing

1. Follow Clean Architecture principles
2. Write tests for new features
3. Update documentation
4. Follow Go coding standards
5. Use conventional commit messages
