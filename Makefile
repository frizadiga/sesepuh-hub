.PHONY: install update tidy dev run clean build ollama openai

BINARY_NAME=sesepuh-hub
# PROMPT=""
# PROMPT="eli5 general relativity"
# PROMPT="write me 50 words haiku"
PROMPT="what model you currently use"

install:
	go mod download

update:
	go get -u

tidy:
	go mod tidy

dev:
	go run -v . --prompt $(PROMPT)

run:
	./$(BINARY_NAME) $(PROMPT)

clean:
	go clean
	rm -f $(BINARY_NAME)

build:
	go build -o $(BINARY_NAME) .

ollama:
	__LLM_MAIN_ENTRY_VENDOR=ollama go run . $(PROMPT)

openai:
	__LLM_MAIN_ENTRY_VENDOR=openai go run . $(PROMPT)
