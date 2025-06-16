package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var MISTRAL_API_KEY = os.Getenv("MISTRAL_API_KEY")
var MISTRAL_MODEL = GetModelToUse("MISTRAL_MODEL", "mistral-small-latest")
var MISTRAL_API_URL = GetEnv("MISTRAL_API_URL", "https://api.mistral.ai/v1/chat/completions")

func ModMistral(prompt *string) {
	if os.Getenv("SESEPUH_HUB_RES_ONLY") != "1" {
		fmt.Printf("\nMistral model: %s\n\n", MISTRAL_MODEL)
	}

	isStreaming := GetEnv("SESEPUH_HUB_STREAMING", "0")

	if isStreaming == "1" {
		ModMistralStream(prompt)
	} else {
		ModMistralSync(prompt)
	}
}

func ModMistralSync(prompt *string) {
	reqBody := map[string]any{
		"model": MISTRAL_MODEL,
		"messages": []map[string]string{
			{"role": "user", "content": *prompt},
		},
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequestWithContext(context.TODO(), "POST", MISTRAL_API_URL, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+MISTRAL_API_KEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println(result.Choices[0].Message.Content)
}

func ModMistralStream(prompt *string) {
	reqBody := map[string]any{
		"model":  MISTRAL_MODEL,
		"stream": true,
		"messages": []map[string]string{
			{"role": "user", "content": *prompt},
		},
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequestWithContext(context.TODO(), "POST", MISTRAL_API_URL, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+MISTRAL_API_KEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // stream ends
		}

		// Only process data lines
		if len(line) < 6 || line[:6] != "data: " {
			continue
		}

		// Remove the "data: " prefix and trim whitespace
		payload := line[6:]
		if payload == "[DONE]\n" || payload == "[DONE]" {
			fmt.Println()
			break
		}

		var chunk struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content,omitempty"`
				} `json:"delta"`
			} `json:"choices"`
		}

		if err := json.Unmarshal([]byte(payload), &chunk); err != nil {
			fmt.Fprintf(os.Stderr, "unmarshal error: %v\n", err)
			continue
		}

		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}
}
