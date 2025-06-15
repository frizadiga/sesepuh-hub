BINARY_NAME=sesepuh-hub
# PROMPT=""
# PROMPT="eli5 general relativity"
# PROMPT="write me 50 words haiku"
PROMPT="what model you currently use"

all: dev

install:
	go mod download

update:
	go get -u

tidy:
	go mod tidy

dev:
	go run -v . --prompt $(PROMPT)

start:
	./$(BINARY_NAME) --prompt $(PROMPT)

clean:
	go clean
	rm -f $(BINARY_NAME)

build:
	go build -o $(BINARY_NAME) .

ollama:
	SESEPUH_HUB_VENDOR=ollama go run . --prompt $(PROMPT)

openai:
	SESEPUH_HUB_VENDOR=openai go run . --prompt $(PROMPT)

gemini:
	SESEPUH_HUB_VENDOR=google go run . --prompt $(PROMPT)

xai:
	SESEPUH_HUB_VENDOR=xai go run . --prompt $(PROMPT)

claude:
	SESEPUH_HUB_VENDOR=anthropic go run . --prompt $(PROMPT)

mistral:
	SESEPUH_HUB_VENDOR=mistral go run . --prompt $(PROMPT)

.PHONY: install update tidy dev run clean build ollama openai gemini xai claude mistral
