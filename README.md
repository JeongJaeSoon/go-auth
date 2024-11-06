# go-auth

go-auth is a simple authentication system built with Go. It includes a basic user model, authentication, and authorization.

# Overview

## Features

- User authentication and authorization
- gRPC and RESTful APIs
- Middleware for handling authentication and authorization
- Session management using Redis
- Database integration with PostgreSQL

## Prerequisites

- Go 1.23.2
- Docker
- Docker Compose (for development)

## Getting Started

1. Clone the repository

```bash
git clone https://github.com/JeongJaeSoon/go-auth.git
```

2. Build the project

```bash
go build -o go-auth ./cmd/server/main.go
```

3. Run the project

```bash
./go-auth
```

# Development

## Setup

```bash
# This section is under preparation and will be provided once ready.
docker compose up --build
```

## Directory Structure

```
go-auth/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── config/
│   ├── config.yaml          # Configuration files
│   └── config.go            # Configuration management
├── internal/
│   ├── auth/
│   │   ├── service.go       # Authentication service logic
│   │   └── handler.go       # HTTP and gRPC handlers
│   ├── db/
│   │   └── db.go            # Database initialization and connection
│   ├── logging/
│   │   └── logger.go        # Logger initialization and configuration
│   ├── middleware/
│   │   └── auth_middleware.go # Authentication middleware
│   ├── proto/
│   │   └── auth.proto       # gRPC protocol definitions
│   └── utils/
│       └── utils.go         # Utility functions
├── pkg/
│   └── models/
│       └── user.go          # User model definitions
├── test/
│   ├── integration/
│   │   └── auth_integration_test.go # Integration tests
│   └── e2e/
│       └── auth_e2e_test.go        # End-to-end tests
├── Dockerfile               # Docker image build configuration
├── docker-compose.yaml      # Docker Compose configuration for development
├── go.mod                   # Go module initialization file
└── go.sum                   # Go module dependency checksum file
```

## Dependency Injection

This project utilizes [Wire](https://github.com/google/wire) for dependency injection. To install Wire, run:

```bash
go install github.com/google/wire/cmd/wire@latest
```
