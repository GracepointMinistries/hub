.PHONY: generate
generate:
	@echo "Running go generate"
	@SWAGGER_GENERATE_EXTENSION=false go generate
