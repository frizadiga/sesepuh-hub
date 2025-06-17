package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var OLLAMA_MODEL = GetModelToUse("OLLAMA_MODEL", "deepseek-coder")
var resultBufOllama strings.Builder

// func ModOllama(prompt *string) {
// 	cmd := exec.Command("ollama", "run", OLLAMA_MODEL, *prompt)
//
// 	output, err := cmd.Output()
// 	if err != nil {
// 		log.Fatal("failed to run ollama: %w", err)
// 	}
//
// 	fmt.Println(string(output))
// 	WriteRespToFile(output, "")
// }

func ModOllama(prompt *string) {
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
			fmt.Print(text)
			resultBufOllama.WriteString(text)
		}
		if err != nil {
			break
		}
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal("ollama command failed: %w", err)
	}

	fmt.Println() // add newline after stream ends
	WriteRespToFile([]byte(resultBufOllama.String()), "")
}
