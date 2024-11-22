include .env

dev: 
	@go run ./cmd/main.go

build: 
	@templ generate
	@go build -o ./bin/${PROJECT_NAME} ./cmd/app/main.go

run: build 
	@./bin/${PROJECT_NAME} 

test:
	@go test ./... -v -count=1 -cover

templ:
	@templ generate --watch --proxy=http://localhost:${HTTP_PORT}

air:
	@air -c .air.toml

air-build:
	@templ generate
	@go build -o ./tmp/${PROJECT_NAME} ./cmd/app/main.go