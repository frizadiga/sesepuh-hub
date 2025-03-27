package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
var __OPENAI_MODEL string = GetEnv("__OPENAI_MODEL", "gpt-4o-mini")

var clientOpenAI = openai.NewClient(
	option.WithAPIKey(OPENAI_API_KEY), // defaults to os.LookupEnv("OPENAI_API_KEY")
)

func ModOpenAI(prompt string) {
	fmt.Printf("\nOpenAI model: %s\n\n", __OPENAI_MODEL)
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
		Model: openai.ChatModelGPT4o,
	})

	if err != nil {
		panic(err.Error())
	}

	println(chatCompletion.Choices[0].Message.Content)
}

func ModOpenAIStream(prompt string) {
	ctx := context.Background()
	stream := clientOpenAI.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Seed:  openai.Int(0),
		Model: openai.ChatModelGPT4o,
	})

	// optionally, an accumulator helper can be used
	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if _, ok := acc.JustFinishedContent(); ok {
			println() // newline after last stream chunk
			// println("Content stream finished:", content)
		}

		// if using tool calls
		if tool, ok := acc.JustFinishedToolCall(); ok {
			println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
		}

		if refusal, ok := acc.JustFinishedRefusal(); ok {
			println("Refusal stream finished:", refusal)
		}

		// it's best to use chunks after handling JustFinished events
		if len(chunk.Choices) > 0 {
			print(chunk.Choices[0].Delta.Content)
		}
	}

	if stream.Err() != nil {
		panic(stream.Err())
	}

	// after the stream is finished, acc can be used like a ChatCompletion
	_ = acc.Choices[0].Message.Content
}
