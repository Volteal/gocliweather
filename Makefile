build:
	@go build -o ./bin/gocliweather ./cmd/gocliweather/main.go

install:
	@go install -v ./cmd/gocliweather

run: 
	@make build
	./bin/gocliweather