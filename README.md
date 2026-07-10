# Bookmark Service API
- This is the backend service to serve APIs related to bookmark

## Installation
### Dependencies
- Run: `go mod download`
### Development tools
- Mockery: `go install github.com/vektra/mockery/v2@v2.52.2`
- Swaggo: `go install github.com/swaggo/swag/cmd/swag@latest`

## Run command
### Local environment
- `make dev-run`
  - Access `http://localhost:8080/swagger/index.html#/` for the API docs
- `make docker-build`
  - build app docker image
- `make test`
  - Run unit tests