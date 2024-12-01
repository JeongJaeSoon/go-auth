# OpenAPI Specification

This directory contains OpenAPI specifications for the authentication service API.

## Directory Structure

```text
api/openapi/
├── README.md           # This file
├── template.yml        # Base template for API specs
├── config.yml         # Common configuration for all APIs
├── health.yml         # Health check API spec
└── health.config.yml  # Health check API specific configuration

tmp/openapi/           # Temporary files (not in version control)
├── merged.yml         # Merged OpenAPI specs
├── merged.config.yml  # Merged configuration
└── *.config.yml       # Temporary configuration files
```

## Requirements

Before you start, make sure you have the following tools installed:

- [oapi-codegen](https://github.com/deepmap/oapi-codegen) - OpenAPI code generator

  ```bash
  go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
  ```

- [yq](https://github.com/mikefarah/yq) - YAML processor

  ```bash
  # macOS
  brew install yq

  # Linux
  sudo wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 -O /usr/bin/yq
  sudo chmod +x /usr/bin/yq
  ```

## Configuration Structure

1. Common Configuration (`config.yml`):

   ```yaml
   generate:
     fiber-server: true
     models: true
     embedded-spec: true
   ```

2. Domain-specific Configuration (e.g., `health.config.yml`):

   ```yaml
   package: health
   output: internal/generated/health/api.gen.go
   ```

The build process will merge these configurations automatically in the `tmp/openapi` directory.

## Adding a New API Specification

To add a new API specification:

1. Create a new YAML file for your API spec (e.g., `auth.yml`):

   ```yaml
   openapi: 3.0.0
   paths:
     /auth/login:
       post:
         summary: User Login
         # ... rest of your API spec
   components:
     schemas:
       LoginRequest:
         # ... your schema definitions
   ```

2. Create a configuration file for your API spec (e.g., `auth.config.yml`):

   ```yaml
   package: auth
   output: internal/generated/auth/api.gen.go
   ```

3. Run code generation:

   ```bash
   make openapi
   ```

This will:

- Generate domain-specific code in `internal/generated/auth/`
- Create temporary files in `tmp/openapi/`
- Generate combined validation code in `internal/generated/api.gen.go`

## Code Organization

The code generation process creates two types of files:

1. Domain-specific files (`internal/generated/<domain>/api.gen.go`):
   - Contains models and interfaces for your specific domain
   - Used by domain handlers for type safety and implementation

2. Merged API spec (`internal/generated/api.gen.go`):
   - Contains combined OpenAPI validation
   - Used by the server for request validation

## Best Practices

1. Keep API specs modular and domain-focused
2. Use common patterns and schemas across specs
3. Document all endpoints and models thoroughly
4. Use semantic versioning for API versions
5. Test generated code before committing
6. Never commit temporary files in `tmp/` directory

## Common Issues

1. **Generation Issues**: If you encounter any issues with code generation:

   ```bash
   make clean
   make openapi
   ```

2. **Type Mismatches**: Ensure consistent types across specs:
   - Use ISO8601 for dates (`date-time` format)
   - Use consistent string formats for IDs
   - Document nullable fields properly
