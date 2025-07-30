# Webstack Template

A modern full-stack web application template built with Go and React.

## Overview

This template provides a solid foundation for building web applications with a Go backend and React frontend. It includes a complete development setup with Docker support, build automation, and modern tooling.

## Key Features

- **Go Backend**: RESTful API server built with Gin framework
- **React Frontend**: Modern React 18 application with TypeScript and Vite
- **WebSocket Support**: Ready for real-time communication between server and client
- **Docker Ready**: Complete containerization setup for development and production
- **Build Automation**: Makefile for common development tasks
- **Configuration Management**: YAML-based configuration with environment variable support

## Architecture

The template uses a clean architecture pattern:

### Components

- **Server**: Go-based HTTP server with embedded static file serving
- **Web UI**: React-based frontend with Tailwind CSS for styling
- **Agent**: Lightweight background process example for distributed architectures

### Technology Stack

**Backend (Go 1.24):**
- **Framework**: Gin for HTTP server and RESTful API
- **Configuration**: Viper for configuration management
- **Static Assets**: Go's embed package to serve frontend files

**Frontend (React + Vite):**
- **Framework**: React 18 with TypeScript
- **Build Tool**: Vite for fast development and builds
- **Styling**: Tailwind CSS for utility-first styling
- **State Management**: React Query and Zustand (configured but minimal)
- **UI Components**: Ready for component library integration

**Development Tools:**
- **Docker**: Multi-stage builds for both server and agent
- **Make**: Build automation and common tasks
- **ESLint**: Code linting for JavaScript/TypeScript

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

### 4. Code Quality

```bash
# Lint everything (Go + Frontend)
make lint

# Lint just Go code
make lint-go

# Lint just frontend code
make lint-web
```

## Configuration

The application uses YAML configuration with environment variable support:

```yaml
# config.yaml
server:
  port: 8080
  host: "0.0.0.0"
  tls:
    enabled: false

agent:
  server_url: "ws://localhost:8080"
  external_service_url: "http://localhost:8080"
  token: ""
```

Environment variables:
- `WEBSTACK_TOKEN` - Authentication token for agent
- `WEBSTACK_SERVER_URL` - WebSocket server URL
- `EXTERNAL_SERVICE_URL` - External service URL (example)

## Development Workflow

### Running the Server
```bash
# Development mode (hot reload with air recommended)
go run cmd/server/main.go

# The server will start on http://localhost:8080
```

### Building with Frontend
```bash
# Build React app first
cd web && npm run build

# This places built files in internal/server/static/
# The server automatically serves the React app

# Start the server
go run cmd/server/main.go
```

### Agent Example
```bash
# The agent requires environment variables:
export WEBSTACK_TOKEN="test-token"
export WEBSTACK_SERVER_URL="ws://localhost:8080"

# Run the agent (lightweight hello world background process)
go run cmd/agent/main.go
```

## Project Structure

```
webstack/
├── cmd/                    # Main applications
│   ├── server/            # Server main entry point
│   └── agent/             # Agent main entry point
├── internal/              # Private application code
│   ├── server/           # Server implementation
│   ├── agent/            # Agent implementation
│   └── common/           # Shared code
├── pkg/                   # Public library code
│   └── config/           # Configuration management
├── web/                   # React frontend
│   ├── src/              # React source code
│   ├── public/           # Static assets
│   └── dist/             # Build output
├── docs/                  # Documentation
├── config.yaml           # Default configuration
├── docker-compose.yml    # Development environment
├── Dockerfile.server     # Server container
├── Dockerfile.agent      # Agent container
└── Makefile              # Build automation
```

## Customization

This template is designed to be easily customized for your specific needs:

1. **Backend**: Modify the server routes and handlers in `internal/server/`
2. **Frontend**: Update the React components in `web/src/`
3. **Configuration**: Adjust settings in `config.yaml` and `pkg/config/`
4. **Docker**: Customize the Dockerfiles for your deployment needs
5. **Build**: Modify the Makefile for your workflow
6. **Agent**: Extend the lightweight agent for your distributed architecture needs

## Security Features

- HTTPS support with TLS configuration
- Environment-based configuration for secrets
- Docker security best practices

## License

This template is provided as-is for educational and development purposes.