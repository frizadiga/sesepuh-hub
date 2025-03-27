package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
)

var __OPENAI_MODEL string = GetEnvWithDefault("__OPENAI_MODEL", "gpt-4o-mini")

var client = openai.NewClient()

func ModOpenAI(prompt string) {
	fmt.Printf("\nOpenAI model: %s\n\n", __OPENAI_MODEL)
	isSesepuhNeedStream := GetEnvWithDefault("SESEPUH_NEED_STREAM", "0")

	if isSesepuhNeedStream == "1" {
		ModOpenAIStream(prompt)
	} else {
		ModOpenAISync(prompt)
	}
}

func ModOpenAISync(prompt string) {
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
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
	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
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
