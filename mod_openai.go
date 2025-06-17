package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
var OPENAI_MODEL = GetModelToUse("OPENAI_MODEL", "gpt-4o-mini")

var clientOpenAI = openai.NewClient(
	option.WithAPIKey(OPENAI_API_KEY), // defaults to os.LookupEnv("OPENAI_API_KEY")
)

var resultBufOpenAI strings.Builder

func ModOpenAI(prompt *string) {
	if os.Getenv("SESEPUH_HUB_RES_ONLY") != "1" {
		fmt.Printf("\nOpenAI model: %s\n\n", OPENAI_MODEL)
	}

	isStreaming := GetEnv("SESEPUH_HUB_STREAMING", "0")

	if isStreaming == "1" {
		ModOpenAIStream(prompt)
	} else {
		ModOpenAISync(prompt)
	}
}

func ModOpenAISync(prompt *string) {
	chatCompletion, err := clientOpenAI.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(*prompt),
		},
		Model: OPENAI_MODEL,
	})

	if err != nil {
		panic(err.Error())
	}

	content := chatCompletion.Choices[0].Message.Content

	// @NOTE: `fmt.Println` ensure write to stdout not stderr, so bash `$()` and other CLI pipe can capture
	fmt.Println(content)

	WriteRespToFile([]byte(content), "")
}

func ModOpenAIStream(prompt *string) {
	ctx := context.Background()
	stream := clientOpenAI.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(*prompt),
		},
		Seed:  openai.Int(0),
		Model: OPENAI_MODEL,
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
			content := chunk.Choices[0].Delta.Content
			fmt.Print(content)
			resultBufOpenAI.WriteString(content)
		}
	}

	if stream.Err() != nil {
		panic(stream.Err())
	}

	// after the stream is finished, acc can be used like a ChatCompletion
	_ = acc.Choices[0].Message.Content

	WriteRespToFile([]byte(resultBufOpenAI.String()), "")
}
