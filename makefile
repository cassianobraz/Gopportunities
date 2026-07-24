PHONY: default run build test docs clean
APP_NAME=gopportunities

default: run-with-docs

run:
	@go run cmd/api/main.go
run-with-docs:
	@swag init -g cmd/api/main.go
	@go run cmd/api/main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	swag init -g cmd/api/main.go
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs