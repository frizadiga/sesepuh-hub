package main

import (
	"fmt"
	"log"
	"os/exec"
)

var __OLLAMA_MODEL = GetEnv("__OLLAMA_MODEL", "deepseek-coder")

func ModOllama(prompt *string) {
	cmd := exec.Command("ollama", "run", __OLLAMA_MODEL, *prompt)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal("failed to run ollama: %w", err)
	}

	fmt.Println(string(output))
}
