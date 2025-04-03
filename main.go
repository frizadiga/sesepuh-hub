package main

import (
	"flag"
	"fmt"
	"os"
)

// read shell env __LLM_MAIN_ENTRY_VENDOR
var vendor string = os.Getenv("__LLM_MAIN_ENTRY_VENDOR")

func main() {
	prompt := getPromptFlag()
	mockRole := "🧙 Sesepuh Hub"
	if os.Getenv("LLM_RES_ONLY") != "1" {
		fmt.Println(mockRole)
	}

	// @NOTE: keep it here for debugging
	// fmt.Println("[DEBUG] vendor:", vendor)
	// fmt.Println("[DEBUG] prompt:", prompt)

	if vendor == "openai" {
		ModOpenAI(prompt)
	}

	if vendor == "ollama" {
		ModOllama(prompt)
	}

	if vendor == "anthropic" {
		ModAnthropic(prompt)
	}

	if vendor == "xai" {
		ModXAI(prompt)
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

// func getPrompt() string {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Error: Please provide a prompt")
// 		os.Exit(1)
// 	}
//
// 	var prompt string = os.Args[1]
// 	// todo something like "$*" in bash
// 	// var prompt string = strings.Join(os.Args[1:], " ")
//
// 	// @NOTE: keep it here for debugging
// 	// print all args
// 	// for i, arg := range os.Args {
// 	// 	fmt.Println(i, arg)
// 	// }
//
// 	return prompt
// }
