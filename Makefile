.PHONY: generate
generate:
	@echo "Running go generate"
	@SWAGGER_GENERATE_EXTENSION=false go generate
	@#echo "Generating dart client"
	@#swagger-codegen generate -i swagger.json -l dart -o dart/client

.PHONY: deps
deps:
	@echo "Installing dependencies"
	@brew tap go-swagger/go-swagger
	@brew install go-swagger swagger-codegen@2
