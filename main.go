package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var vendor string = os.Getenv("SESEPUH_HUB_VENDOR")

func main() {
	prompt := getPromptFlag()
	mockRole := "🧙 Sesepuh Hub"
	if os.Getenv("LLM_RES_ONLY") != "1" {
		fmt.Println(mockRole)
	}

	vendorHandlers := map[string]func(*string){
		"openai":    ModOpenAI,
		"ollama":    ModOllama,
		"anthropic": ModAnthropic,
		"xai":       ModXAI,
		"google":    ModGoogle,
		"mistral":   ModMistral,
	}

	if handler, exists := vendorHandlers[vendor]; exists {
		handler(&prompt)
	} else {
		log.Fatalf("Error: Unknown vendor '%s'\n", vendor)
	}
}

func getPromptFlag() string {
	promptFlag := flag.String("prompt", "", "The prompt to use")
	flag.Parse()

	if *promptFlag == "" {
		log.Fatal("Error: Please provide a prompt using --prompt flag")
	}
	return *promptFlag
}
