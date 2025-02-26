package main

import (
	"fmt"
	"os"
	// "log"
	// "net/http"
)

// read shell env __LLM_MAIN_ENTRY_VENDOR
var vendor string = os.Getenv("__LLM_MAIN_ENTRY_VENDOR")

func main() {
	prompt := getPrompt()
	fmt.Println("vendor:", vendor)
	fmt.Println("prompt:", prompt)

	if vendor == "openai" {
		ModOpenAI(prompt)
	}

	if vendor == "ollama" {
		ModOllama(prompt)
	}
}

func getPrompt() string {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a prompt")
		os.Exit(1)
	}

	// get first "actual" arg from cli
	var prompt string = os.Args[1]

	// print all args
	// for i, arg := range os.Args {
	// 	fmt.Println(i, arg)
	// }

	return prompt
}
