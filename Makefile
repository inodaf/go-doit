# Tasks
.PHONY: all
all: ./bin/todo

.PHONY: clean
clean:
	@rm -rf ./bin

# Recipes
./bin/todo : $(shell find . -name '*.go')
	@go build -o ./bin/todo ./cmd/todo
