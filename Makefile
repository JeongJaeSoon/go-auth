# Makefile

# Generate wire_gen.go if it does not exist or wire.go has changed
generate:
	wire

# Generate OpenAPI code
openapi:
	mkdir -p internal/generated
	oapi-codegen -config api/openapi/config.yaml api/openapi/health.yml

# Run the application, ensuring wire_gen.go is included
run: generate openapi
	go run main.go wire_gen.go

# Clean up generated files
clean:
	rm -f wire_gen.go
	rm -rf internal/generated
