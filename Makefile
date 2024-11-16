# Makefile

# Generate wire_gen.go if it does not exist or wire.go has changed
generate:
	wire

# Run the application, ensuring wire_gen.go is included
run: generate
	go run main.go wire_gen.go

# Clean up wire_gen.go if needed
clean:
	rm -f wire_gen.go
