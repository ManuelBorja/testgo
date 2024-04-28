build:
	@go build -o bin/testgo

run: build
	@./bin/testgo

test:
	@go test -v ./...
