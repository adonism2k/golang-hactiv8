APP_BINARY=final-project
TEST_BINARY=test-project

up:
	@echo "Stopping docker images (if running...)"
	cd ./docker && docker compose down --volumes
	@echo "Building docs folder for swagger API documentation..."
	swag init -g ./cmd/final-project/main.go
	@echo "Building go binaries..."
	env GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o ${APP_BINARY} ./cmd/final-project
	@echo "Starting docker containers..."
	cd ./docker && docker compose rm -fsv && docker compose up --build

down:
	@echo "Stopping docker containers..."
	cd ./docker && docker compose down --volumes

test:
	@echo "Building go test binaries..."
	env GOOS=linux GOARCH=386 CGO_ENABLED=0 go test -c ${TEST_BINARY}
	@echo "Starting docker containers..."
	cd ./docker && docker compose up test && docker compose rm -fsv

binary:
	@echo "Building go binaries..."
	env GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o ${APP_BINARY} ./cmd/final-project

test-binary:
	@echo "Building go test binaries..."
	env GOOS=linux GOARCH=386 go test -c ${TEST_BINARY}

clean:
	@echo "Cleaing docker containers..."
	cd ./docker && docker compose rm -fsv && docker-compose down --volumes