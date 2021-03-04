.PHONY: generate
generate: clean-generate
	@echo "Running go generate"
	@SWAGGER_GENERATE_EXTENSION=false go generate
	@echo "Generating swagger clients"
	@swagger-codegen generate -i swagger.json -l dart -o dart/client
	@swagger-codegen generate -i swagger.json -l go -o client -DpackageName=client
	@echo "# This folder is generated" > models/README.md
	@echo "# This folder is generated" > dart/client/README.md
	@echo "# This folder is generated" > client/README.md

.PHONY: clean-generate
clean-generate:
	@echo "Cleaning up generated files"
	@rm -rf dart/client client

.PHONY: deps
deps:
	@echo "Installing dependencies"
	@brew tap go-swagger/go-swagger
	@brew install go-swagger swagger-codegen@2
