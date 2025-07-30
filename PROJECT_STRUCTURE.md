# Webstack Template Project Structure

This document describes the foundational project structure for the Webstack template.

## Directory Structure

```
webstack/
├── cmd/                    # Main applications
│   ├── server/            # Server main entry point
│   └── agent/             # Agent main entry point
├── internal/              # Private application code
│   ├── server/           # Server implementation
│   ├── agent/            # Agent implementation
│   └── common/           # Shared code between server and agent
├── pkg/                   # Public library code
│   ├── config/           # Configuration management
│   ├── auth/             # Authentication utilities (placeholder)
│   └── tunnel/           # Communication protocols (placeholder)
├── web/                   # React frontend
│   ├── src/              # React source code
│   ├── public/           # Static assets
│   └── dist/             # Build output (goes to internal/server/static)
├── docs/                  # Documentation
├── config.yaml           # Default configuration
├── docker-compose.yml    # Development environment
├── Dockerfile.server     # Server container
├── Dockerfile.agent      # Agent container
└── Makefile              # Build automation
```

## Technology Stack

### Backend (Go 1.24)
- **Framework**: Gin for HTTP server and API
- **Configuration**: spf13/viper for config management
- **Authentication**: Ready for golang-jwt/jwt token handling
- **Static Assets**: Go embed for serving frontend

### Frontend (React + Vite)
- **Framework**: React 18 with TypeScript
- **Build Tool**: Vite for fast development
- **Styling**: Tailwind CSS
- **State Management**: @tanstack/react-query + zustand (configured)
- **Terminal**: xterm.js (ready for console features)

### Agent (Minimal Go)
- **Ultra-lightweight**: Standard library focus
- **Background Process**: Simple hello world daemon
- **Configurable**: Environment-based configuration

## Quick Start

### 1. Development Setup

```bash
# Install Go dependencies
go mod tidy

# Install frontend dependencies
cd web && npm install

# Start development server
make dev-server

# In another terminal, start frontend dev server
make dev-web
```

### 2. Production Build

```bash
# Build everything
make build

# Or build individual components
make build-web
make build-server
make build-agent
```

### 3. Docker Development

```bash
# Build Docker images
make docker-build

# Start full environment
make docker-up

# View logs
make docker-logs

# Stop environment
make docker-down
```

## Current Implementation Status

✅ **Completed Foundation Components:**

1. **Project Structure**: Standard Go project layout with cmd/, internal/, pkg/
2. **Go Server**: Gin server with health endpoints and embedded static file serving
3. **React Frontend**: Fully working Vite-based React app with Tailwind CSS and Hello World page
4. **Agent Structure**: Lightweight hello world background process
5. **Configuration**: Viper-based config management
6. **Docker Setup**: Development environment with Docker Compose
7. **Build System**: Makefile for common development tasks
8. **Static File Integration**: Embedded file serving for React assets

🚧 **Template Extension Points:**

1. **Authentication System**: JWT token implementation ready for extension
2. **API Endpoints**: RESTful API structure ready for business logic
3. **Real-time Features**: WebSocket framework ready for implementation
4. **Database Integration**: Ready for ORM or database drivers
5. **Frontend Components**: Component library integration point

## Development Workflow

### Running the Server
```bash
# Development mode (serves React app)
go run cmd/server/main.go

# The server will start on http://localhost:8080
# Visit the URL to see the Hello World page
```

### Building with Frontend
```bash
# Build React app first
cd web && npm run build

# This places the built files in internal/server/static/
# The server will automatically serve the React app

# Start the server
go run cmd/server/main.go
```

### Agent Testing
```bash
# The agent requires environment variables:
export WEBSTACK_TOKEN="test-token"
export WEBSTACK_SERVER_URL="ws://localhost:8080"

# Run the agent (lightweight hello world background process)
go run cmd/agent/main.go
```

## Architecture Patterns

This template implements several architectural patterns:

- **Clean Architecture**: Separation of concerns with internal/pkg structure
- **Embedded Assets**: Self-contained binary with frontend included
- **Background Processing**: Lightweight agent pattern for distributed systems
- **Docker-First**: Container-ready development and deployment
- **Configuration-Driven**: YAML + environment variable configuration

## Extension Guidelines

### Adding New API Endpoints
1. Add routes in `internal/server/server.go`
2. Implement handlers following the existing pattern
3. Update configuration if needed

### Frontend Development
1. Add components in `web/src/components/`
2. Use Tailwind CSS for styling
3. Leverage React Query for API state management

### Agent Features
1. Extend the agent functionality in `internal/agent/`
2. Add configuration options in `pkg/config/`
3. Follow the existing lightweight pattern

### Database Integration
1. Add database configuration to `pkg/config/`
2. Implement repositories in `internal/`
3. Update Docker Compose for database services

## Build and Deployment

The template supports multiple deployment strategies:

- **Single Binary**: Embedded frontend with Go binary
- **Container Deployment**: Docker images for server and agent
- **Development Mode**: Vite dev server + Go server
- **Static Deployment**: Build frontend for CDN deployment

This structure provides a solid foundation for building modern web applications with Go and React while maintaining flexibility for different deployment scenarios and architectural patterns.