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

release:
	go build -o $(BINARY_NAME) -ldflags="-s -w" .

ollama:
	SESEPUH_HUB_VENDOR=ollama SESEPUH_HUB_MODEL='' go run . --prompt $(PROMPT)

openai:
	SESEPUH_HUB_VENDOR=openai SESEPUH_HUB_MODEL='' go run . --prompt $(PROMPT)

gemini:
	SESEPUH_HUB_VENDOR=google SESEPUH_HUB_MODEL='' go run . --prompt $(PROMPT)

xai:
	SESEPUH_HUB_VENDOR=xai SESEPUH_HUB_MODEL='' go run . --prompt $(PROMPT)

claude:
	SESEPUH_HUB_VENDOR=anthropic SESEPUH_HUB_MODEL='' go run . --prompt $(PROMPT)

mistral:
	SESEPUH_HUB_VENDOR=mistral SESEPUH_HUB_MODEL='' go run . --prompt $(PROMPT)

.PHONY: all install update tidy dev start clean build release ollama openai gemini xai claude mistral
