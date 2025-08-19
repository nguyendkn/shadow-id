# Shadow ID Makefile

.PHONY: help dev build test clean install deps lint fmt vet

# Default target
help: ## Show this help message
	@echo "Shadow ID - Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Development
dev: ## Start development server with hot reload
	wails dev

build: ## Build the application
	wails build

build-prod: ## Build production version with optimizations
	wails build -clean -upx

# Testing
test: ## Run all tests
	go test ./...

test-verbose: ## Run tests with verbose output
	go test -v ./...

test-coverage: ## Run tests with coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Code quality
lint: ## Run linter
	golangci-lint run

fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

# Dependencies
deps: ## Download and tidy dependencies
	go mod download
	go mod tidy

install: ## Install development tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Cleanup
clean: ## Clean build artifacts
	rm -rf build/bin/
	rm -rf frontend/dist/
	rm -rf frontend/.next/
	rm -f coverage.out coverage.html

# Frontend
frontend-install: ## Install frontend dependencies
	cd frontend && npm install

frontend-build: ## Build frontend
	cd frontend && npm run build

frontend-dev: ## Start frontend development server
	cd frontend && npm run dev

# Database (for future use)
db-migrate: ## Run database migrations
	@echo "Database migrations not implemented yet"

db-seed: ## Seed database with test data
	@echo "Database seeding not implemented yet"

# Docker (for future use)
docker-build: ## Build Docker image
	docker build -t shadow-id .

docker-run: ## Run Docker container
	docker run -p 8080:8080 shadow-id

# Generate
generate: ## Generate code (mocks, etc.)
	go generate ./...

# All-in-one commands
setup: install deps frontend-install ## Setup development environment

check: fmt vet lint test ## Run all checks

build-all: clean frontend-build build ## Build everything

# Environment
env-example: ## Create example environment file
	@echo "Creating .env.example..."
	@cat > .env.example << 'EOF'
# Application Configuration
APP_NAME=shadow-id
APP_VERSION=1.0.0
APP_ENV=development

# Logging
LOG_LEVEL=info

# Database (for future use)
DB_DRIVER=memory
DB_HOST=localhost
DB_PORT=5432
DB_NAME=shadow_id
DB_USER=postgres
DB_PASSWORD=
DB_SSL_MODE=disable

# Server (for future use)
SERVER_HOST=localhost
SERVER_PORT=8080
EOF
	@echo ".env.example created successfully"
