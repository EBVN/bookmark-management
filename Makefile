.PHONY: run swagger test dev-run

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go

COVERAGE_EXCLUDE=mocks|main.go|test|docs

test:
	go test ./... -coverpkg=./... -covermode=atomic -p 1 -coverprofile=coverage.tmp
	grep -vE "$(COVERAGE_EXCLUDE)" coverage.tmp > coverage.out
	go tool cover -html=coverage.out -o coverage.html

docker-build:
	docker build -t ebvn_test:latest .

dev-run: swagger run