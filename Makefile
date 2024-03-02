build:
	@go build -o ./bin/cli ./cmd/cli

run %: build
	@./bin/cli %