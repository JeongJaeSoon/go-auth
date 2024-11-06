# go-auth

go-auth is a simple authentication system built with Go. It includes a basic user model, authentication, and authorization.

## Directory Structure
```
go-auth/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── auth/
│   │   ├── service.go       # Authentication service logic
│   │   └── handler.go       # HTTP and gRPC handlers
│   ├── db/
│   │   └── db.go            # Database initialization and connection
│   ├── logging/
│   │   ├── logger.go        # Logger initialization and configuration
│   │   └── config.go        # Logging configuration management
│   ├── proto/
│   │   └── auth.proto       # gRPC protocol definitions
│   ├── middleware/
│   │   └── auth_middleware.go # Authentication middleware
│   └── utils/
│       └── utils.go         # Utility functions
├── pkg/
│   └── models/
│       └── user.go          # User model definitions
├── config/
│   └── config.yaml          # Configuration files
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
