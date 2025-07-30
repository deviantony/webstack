.PHONY: build build-server build-agent build-web dev test clean docker-build docker-up docker-down lint lint-go lint-web

# Default target
all: build

# Build everything
build: build-web build-server build-agent

# Build server binary
build-server:
	@echo "Building server..."
	go build -o bin/server ./cmd/server

# Build agent binary
build-agent:
	@echo "Building agent..."
	go build -o bin/agent ./cmd/agent

# Build web frontend
build-web:
	@echo "Building web frontend..."
	cd web && npm install && npm run build

# Development mode (requires separate terminals)
dev-server:
	@echo "Starting development server..."
	go run ./cmd/server

dev-web:
	@echo "Starting development web server..."
	cd web && npm run dev

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf web/dist/
	rm -rf internal/server/static/

# Docker commands
docker-build:
	@echo "Building Docker images..."
	docker build -f Dockerfile.server -t webstack-server .
	docker build -f Dockerfile.agent -t webstack-agent .

docker-up:
	@echo "Starting Docker services..."
	docker-compose up -d

docker-down:
	@echo "Stopping Docker services..."
	docker-compose down

docker-logs:
	docker-compose logs -f

# Initialize Go modules
mod-tidy:
	go mod tidy

# Install web dependencies
web-install:
	cd web && npm install

# Linting commands
lint: lint-go lint-web

lint-go:
	@echo "Linting Go code..."
	go vet ./...
	go fmt ./...
	@if command -v golangci-lint >/dev/null 2>&1; then \
		echo "Running golangci-lint..."; \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

lint-web:
	@echo "Linting frontend code..."
	cd web && npm run lint