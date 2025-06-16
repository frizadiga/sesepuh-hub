package main

import (
	"fmt"
	"log"
	"os/exec"
)

var OLLAMA_MODEL = GetModelToUse("OLLAMA_MODEL", "deepseek-coder")

func ModOllama(prompt *string) {
	cmd := exec.Command("ollama", "run", OLLAMA_MODEL, *prompt)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal("failed to run ollama: %w", err)
	}

	fmt.Println(string(output))
}
