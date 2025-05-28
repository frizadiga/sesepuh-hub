package main

import (
	"fmt"
	"os/exec"
)

var __OLLAMA_MODEL = GetEnv("__OLLAMA_MODEL", "deepseek-coder")

func ModOllama(prompt *string) error {
	cmd := exec.Command("ollama", "run", __OLLAMA_MODEL, *prompt)

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run ollama: %w", err)
	}

	fmt.Println(string(output))
	return nil
}
