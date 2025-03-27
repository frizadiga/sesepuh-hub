.PHONY: build run clean

BINARY_NAME=sesepuh-gpt
# ARGS=""
# ARGS="eli5 general relativity"
# ARGS="write me 50 words haiku"
ARGS="what model you currently use"

install:
	go mod download

update:
	go get -u

tidy:
	go mod tidy

dev:
	go run -v . $(ARGS)

run:
	./$(BINARY_NAME) $(ARGS)

clean:
	go clean
	rm -f $(BINARY_NAME)

build:
	go build -o $(BINARY_NAME) .

ollama:
	__LLM_MAIN_ENTRY_VENDOR=ollama go run . $(ARGS)

openai:
	__LLM_MAIN_ENTRY_VENDOR=openai go run . $(ARGS)
