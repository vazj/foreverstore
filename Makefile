build: 
	@go build -o bin/fs

run: build
	@./bin/fs

test: build
	@go test -v ./...