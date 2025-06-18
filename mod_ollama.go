package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var OLLAMA_MODEL = GetModelToUse("OLLAMA_MODEL", "deepseek-coder")
var resultBufOllama strings.Builder

func ModOllama(prompt *string) {
	if os.Getenv("SESEPUH_HUB_RES_ONLY") != "1" {
		fmt.Printf("\nOllama model: %s\n\n", OLLAMA_MODEL)
	}

	isStreaming := GetEnv("SESEPUH_HUB_STREAMING", "0")

	if isStreaming == "1" {
		ModOllamaStream(prompt)
	} else {
		ModOllamaSync(prompt)
	}
}

func ModOllamaSync(prompt *string) {
	cmd := exec.Command("ollama", "run", OLLAMA_MODEL, *prompt)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal("failed to run ollama: %w", err)
	}

	fmt.Print(string(output))
	WriteRespToFile(output, "")
}

func ModOllamaStream(prompt *string) {
	cmd := exec.Command("ollama", "run", OLLAMA_MODEL, *prompt)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("failed to start ollama: %w", err)
	}

	buffer := make([]byte, 1024)
	for {
		n, err := stdout.Read(buffer)
		if n > 0 {
			text := string(buffer[:n])
			if text != "\n" && text != "" {
				fmt.Print(text)
				resultBufOllama.WriteString(text)
			}
		}
		if err != nil {
			break
		}
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal("ollama command failed: %w", err)
	}

	WriteRespToFile([]byte(resultBufOllama.String()), "")
}
