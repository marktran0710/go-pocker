build:
	@go build -o bin/go-pocker

run: build
	@./bin/go-pocker

test:
	go test -v ./...