build: internal/
	@go build -o ./bin/todo ./cmd/todo

# .PHONY: run
# run %: build
# 	@./bin/cli %