package main

import (
	"flag"
	"fmt"
	"os"
)

var vendor string = os.Getenv("__LLM_MAIN_ENTRY_VENDOR")

func main() {
	prompt := getPromptFlag()
	mockRole := "🧙 Sesepuh Hub"
	if os.Getenv("LLM_RES_ONLY") != "1" {
		fmt.Println(mockRole)
	}

	if vendor == "openai" {
		ModOpenAI(&prompt)
	}

	if vendor == "ollama" {
		ModOllama(&prompt)
	}

	if vendor == "anthropic" {
		ModAnthropic(&prompt)
	}

	if vendor == "xai" {
		ModXAI(&prompt)
	}

	if vendor == "google" {
		ModGoogle(&prompt)
	}

	if vendor == "mistral" {
		ModMistral(&prompt)
	}
}

func getPromptFlag() string {
	promptFlag := flag.String("prompt", "", "The prompt to use")
	flag.Parse()

	if *promptFlag == "" {
		fmt.Println("Error: Please provide a prompt using --prompt flag")
		os.Exit(1)
	}

	return *promptFlag
}
