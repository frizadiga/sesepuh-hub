package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"os"
)

func ModOpenAI(prompt string) {
	fmt.Println("-> ModOpenAI")
	fmt.Println("vendor:", vendor)

	// openai
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(
		option.WithAPIKey(openaiAPIKey), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)
}
