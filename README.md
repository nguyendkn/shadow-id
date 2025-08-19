# Shadow ID

A modern desktop application built with Wails v2, Go backend, and Next.js frontend, following Clean Architecture principles.

## Architecture

This project implements Clean Architecture with clear separation of concerns:

- **Domain Layer**: Pure business logic and entities
- **Application Layer**: Use cases and application services (CQRS pattern)
- **Infrastructure Layer**: External integrations (Wails, storage, etc.)

See [Architecture Documentation](docs/ARCHITECTURE.md) for detailed information.

## Features

- 🏗️ **Clean Architecture**: Maintainable and testable codebase
- 🎯 **CQRS Pattern**: Separate command and query operations
- 🔧 **Dependency Injection**: Loose coupling and easy testing
- 📝 **Structured Logging**: Comprehensive logging system
- 🚨 **Error Handling**: Custom error types with context
- 🧪 **Unit Testing**: Comprehensive test coverage
- 🎨 **Modern UI**: Next.js with Tailwind CSS
- 📱 **Cross-platform**: Windows, macOS, and Linux support

## Project Structure

```
shadow-id/
├── internal/               # Private application code
│   ├── app/               # Application layer (CQRS)
│   ├── domain/            # Domain layer (entities, interfaces)
│   └── infra/             # Infrastructure layer (Wails, storage)
├── pkg/                   # Public packages (utilities)
├── configs/               # Configuration files
├── docs/                  # Documentation
├── frontend/              # Next.js frontend
└── build/                 # Build artifacts
```

## Quick Start

### Prerequisites

- Go 1.21 or later
- Node.js 18 or later
- Wails CLI v2

### Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd shadow-id
   ```

2. **Install dependencies**

   ```bash
   make setup
   # or manually:
   go mod tidy
   cd frontend && npm install
   ```

3. **Run in development mode**

   ```bash
   make dev
   # or:
   wails dev
   ```

4. **Build for production**
   ```bash
   make build
   # or:
   wails build
   ```

## Development

### Available Commands

```bash
make help          # Show all available commands
make dev           # Start development server
make build         # Build application
make test          # Run tests
make lint          # Run linter
make clean         # Clean build artifacts
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific package tests
go test ./internal/domain/entities
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Run vet
make vet

# Run all checks
make check
```

## Configuration

The application uses environment variables for configuration. Copy `.env.example` to `.env` and modify as needed:

```bash
cp .env.example .env
```

Key configuration options:

- `APP_ENV`: Environment (development, production, test)
- `LOG_LEVEL`: Logging level (debug, info, warn, error)
- `DB_DRIVER`: Database driver (memory, postgres)

## API Methods

The application exposes the following methods to the frontend:

### User Management

- `CreateUser(name, email string)`: Create a new user
- `GetUser(id string)`: Get user by ID
- `GetAppInfo()`: Get application information

## Contributing

1. Follow Clean Architecture principles
2. Write tests for new features
3. Update documentation
4. Follow Go coding standards
5. Use conventional commit messages

## License

This project is licensed under the MIT License.
