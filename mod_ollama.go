package main

import (
	"fmt"
	"os/exec"
)

func ModOllama(prompt string) error {
	final_model := "deepseek-coder"
	cmd := exec.Command("ollama", "run", final_model, prompt)

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run ollama: %w", err)
	}

	fmt.Println(string(output))
	return nil
}
