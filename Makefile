.PHONY: generate
generate:
	@echo "Running go generate"
	@SWAGGER_GENERATE_EXTENSION=false go generate
	@echo "Generating swagger clients"
	@swagger-codegen generate -i swagger.json -l dart -o dart/client
	@swagger-codegen generate -i swagger.json -l go -o client -DpackageName=client
	@swagger-codegen generate -i swagger.json -l typescript-fetch -o typescript/client

.PHONY: deps
deps:
	@echo "Installing dependencies"
	@brew tap go-swagger/go-swagger
	@brew install go-swagger swagger-codegen@2
