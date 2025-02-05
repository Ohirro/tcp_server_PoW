.PHONY: all build-server build-client run-server run-client clean

all: build-server build-client

build-server:
	@echo "Building server..."
	@cd server && go build -o server .

build-client:
	@echo "Building client..."
	@cd client && go build -o client .

run-server: build-server
	@echo "Running server..."
	@cd server && ./server

run-client: build-client
	@echo "Running client..."
	@cd client && ./client

clean:
	@echo "Cleaning binaries..."
	@rm -f server/server
	@rm -f client/client
