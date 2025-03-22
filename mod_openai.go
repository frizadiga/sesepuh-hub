package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"os"
)

var __OPENAI_MODEL = getEnv("__OPENAI_MODEL", "o3-mini")

func getModelToUse() string {
	const DEF_MODEL = "o3-mini"
	modelName := DEF_MODEL

	if envModel := __OPENAI_MODEL; envModel != "" {
		modelName = envModel
	}

	return modelName
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func ModOpenAI(prompt string) {
	fmt.Printf("\nOpenAI model: %s\n\n", __OPENAI_MODEL)
	isSesepuhNeedStream := getEnv("SESEPUH_NEED_STREAM", "0")

	if isSesepuhNeedStream == "1" {
		ModOpenAIStream(prompt)
	} else {
		ModOpenAISync(prompt)
	}
}

func ModOpenAISync(prompt string) {
	// openai
	openaiAPIKey := getEnv("OPENAI_API_KEY", "")

	client := openai.NewClient(
		option.WithAPIKey(openaiAPIKey), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}), Model: openai.F(getModelToUse()),
	})

	if err != nil {
		panic(err.Error())
	}

	println(chatCompletion.Choices[0].Message.Content)
}

func ModOpenAIStream(prompt string) {
	client := openai.NewClient()
	ctx := context.Background()

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(getModelToUse()),
	})

	for stream.Next() {
		evt := stream.Current()
		if len(evt.Choices) > 0 {
			print(evt.Choices[0].Delta.Content)
		}
	}

	println() // @NOTE: new line needed due to during stream, the cursor is not moved to new line

	if err := stream.Err(); err != nil {
		panic(err)
	}
}
