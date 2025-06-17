package main

import (
	"context"
	"fmt"
	"os"

	// @NOTE: yes xAI use openai interface
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var XAI_API_KEY = os.Getenv("XAI_API_KEY")
var XAI_MODEL = GetModelToUse("XAI_MODEL", "grok-2-latest")
var XAI_URL = GetEnv("XAI_URL", "https://api.x.ai/v1")

var clientXAI = openai.NewClient(
	option.WithAPIKey(XAI_API_KEY),
	option.WithBaseURL(XAI_URL),
)

func ModXAI(prompt *string) {
	if os.Getenv("SESEPUH_HUB_RES_ONLY") != "1" {
		fmt.Printf("\nXAI model: %s\n\n", XAI_MODEL)
	}

	isStreaming := GetEnv("SESEPUH_HUB_STREAMING", "0")

	if isStreaming == "1" {
		ModXAIStream(prompt)
	} else {
		ModXAISync(prompt)
	}
}

func ModXAISync(prompt *string) {
	chatCompletion, err := clientXAI.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(*prompt),
		},
		Model: XAI_MODEL,
	})

	if err != nil {
		panic(err.Error())
	}

	content := chatCompletion.Choices[0].Message.Content
	fmt.Println(content)

	WriteRespToFile([]byte(content), "")
}

func ModXAIStream(prompt *string) {
	ctx := context.Background()
	stream := clientXAI.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(*prompt),
		},
		Seed:  openai.Int(1),
		Model: XAI_MODEL,
	})

	// optionally, an accumulator helper can be used
	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if _, ok := acc.JustFinishedContent(); ok {
			fmt.Println() // newline after last stream chunk
		}

		// if using tool calls
		if tool, ok := acc.JustFinishedToolCall(); ok {
			fmt.Println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
		}

		if refusal, ok := acc.JustFinishedRefusal(); ok {
			fmt.Println("Refusal stream finished:", refusal)
		}

		// it's best to use chunks after handling JustFinished events
		if len(chunk.Choices) > 0 {
			content := chunk.Choices[0].Delta.Content
			fmt.Print(content)
		}
	}

	if stream.Err() != nil {
		panic(stream.Err())
	}

	// after the stream is finished, acc can be used like a ChatCompletion
	fullContent := acc.Choices[0].Message.Content

	WriteRespToFile([]byte(string(fullContent)), "")
}
