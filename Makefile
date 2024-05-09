build: internal/
	@go build -o ./bin/cli ./cmd/cli

# .PHONY: run
# run %: build
# 	@./bin/cli %