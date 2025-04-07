package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
var __OPENAI_MODEL = GetEnv("__OPENAI_MODEL", "gpt-4o-mini")

var clientOpenAI = openai.NewClient(
	option.WithAPIKey(OPENAI_API_KEY), // defaults to os.LookupEnv("OPENAI_API_KEY")
)

func ModOpenAI(prompt string) {
	if os.Getenv("LLM_RES_ONLY") != "1" {
		fmt.Printf("\nOpenAI model: %s\n\n", __OPENAI_MODEL)
	}

	isSesepuhNeedStream := GetEnv("SESEPUH_NEED_STREAM", "0")

	if isSesepuhNeedStream == "1" {
		ModOpenAIStream(prompt)
	} else {
		ModOpenAISync(prompt)
	}
}

func ModOpenAISync(prompt string) {
	chatCompletion, err := clientOpenAI.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: __OPENAI_MODEL,
		// Model: openai.ChatModelGPT4o,
	})

	if err != nil {
		panic(err.Error())
	}

	// @NOTE: `fmt.Println` ensure write to stdout not stderr, so bash `$()` and other CLI pipe can capture
	fmt.Println(chatCompletion.Choices[0].Message.Content)
}

func ModOpenAIStream(prompt string) {
	ctx := context.Background()
	stream := clientOpenAI.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Seed:  openai.Int(0),
		Model: __OPENAI_MODEL,
		// NOTE: default Model: openai.ChatModelGPT4o,
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
			// @NOTE: `fmt.Print` ensure write to stdout not stderr, so bash `$()` and other CLI pipe can capture
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}

	if stream.Err() != nil {
		panic(stream.Err())
	}

	// after the stream is finished, acc can be used like a ChatCompletion
	_ = acc.Choices[0].Message.Content
}
