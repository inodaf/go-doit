# Tasks
.PHONY: all
all : ./bin/todo

.PHONY: clean
clean :
	@rm -rf ./bin

# Recipes
./bin/todo: $(shell find internal/**/*.go)
	@go build -o ./bin/todo ./cmd/todo
