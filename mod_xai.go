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
// model: "grok-2-latest",
var __XAI_MODEL = GetEnv("__XAI_MODEL", "grok-2-latest")

var clientXAI = openai.NewClient(
	option.WithAPIKey(XAI_API_KEY),
	option.WithBaseURL("https://api.x.ai/v1"),
)

func ModXAI(prompt string) {
	if os.Getenv("LLM_RES_ONLY") != "1" {
		fmt.Printf("\nXAI model: %s\n\n", __XAI_MODEL)
	}

	isSesepuhNeedStream := GetEnv("SESEPUH_NEED_STREAM", "0")

	if isSesepuhNeedStream == "1" {
		ModXAIStream(prompt)
	} else {
		ModXAISync(prompt)
	}
}

func ModXAISync(prompt string) {
	chatCompletion, err := clientXAI.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: __XAI_MODEL,
		// Model: openai.ChatModelGPT4o,
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(chatCompletion.Choices[0].Message.Content)
}

func ModXAIStream(prompt string) {
	ctx := context.Background()
	stream := clientXAI.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Seed:  openai.Int(1),
		Model: __XAI_MODEL,
		// Model: openai.ChatModelGPT4o,
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
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}

	if stream.Err() != nil {
		panic(stream.Err())
	}

	// after the stream is finished, acc can be used like a ChatCompletion
	_ = acc.Choices[0].Message.Content
}
