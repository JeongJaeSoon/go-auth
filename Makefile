# Makefile

.PHONY: generate.wire generate.openapi run clean

generate: generate.wire generate.openapi
# Generate dependency injection code
generate.wire:
	wire

# Generate OpenAPI code for each API spec and merge them
generate.openapi:
	@echo "Creating temporary directory..."
	@mkdir -p tmp/openapi
	@echo "Generating individual API specs..."
	@for spec in api/openapi/*.yml; do \
		if [ "$${spec}" != "api/openapi/config.yml" ] && [ "$${spec}" != "api/openapi/template.yml" ]; then \
			name=$$(basename $${spec} .yml); \
			if [ -f "api/openapi/$${name}.config.yml" ]; then \
				echo "Generating code for $${name}..."; \
				mkdir -p internal/generated/$${name}; \
				yq eval-all '. as $$item ireduce ({}; . * $$item )' api/openapi/config.yml api/openapi/$${name}.config.yml > tmp/openapi/$${name}.config.yml; \
				oapi-codegen -config tmp/openapi/$${name}.config.yml $${spec}; \
			fi \
		fi \
	done
	@echo "Merging OpenAPI specs..."
	@cp api/openapi/template.yml tmp/openapi/merged.yml
	@for spec in api/openapi/*.yml; do \
		if [ "$${spec}" != "api/openapi/config.yml" ] && [ "$${spec}" != "api/openapi/template.yml" ]; then \
			echo "Merging $${spec}..."; \
			yq eval-all '. as $$item ireduce ({}; . * $$item )' tmp/openapi/merged.yml $${spec} > tmp/openapi/merged.tmp.yml && mv tmp/openapi/merged.tmp.yml tmp/openapi/merged.yml; \
		fi \
	done
	@echo "Generating merged API spec..."
	@mkdir -p internal/generated
	@yq eval-all '. as $$item ireduce ({}; . * $$item )' api/openapi/config.yml - > tmp/openapi/merged.config.yml <<< '{"package": "generated", "output": "internal/generated/api.gen.go"}'
	@oapi-codegen -config tmp/openapi/merged.config.yml tmp/openapi/merged.yml

# Run the application
run: generate.wire generate.openapi
	go run main.go wire_gen.go

# Run the application with hot reloading
run.hot: generate.wire generate.openapi
	air

# Clean up generated files
clean:
	rm -f wire_gen.go
	rm -rf internal/generated
	rm -rf tmp/openapi
