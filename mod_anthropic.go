package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go" // imported as anthropic
	"github.com/anthropics/anthropic-sdk-go/option"
)

var ANTHROPIC_API_KEY = os.Getenv("ANTHROPIC_API_KEY")
var ANTHROPIC_MODEL = GetModelToUse("ANTHROPIC_MODEL", "claude-sonnet-4-20250514")

var clientAnthropic = anthropic.NewClient(
	option.WithAPIKey(ANTHROPIC_API_KEY), // defaults to os.LookupEnv("ANTHROPIC_API_KEY")
)

func ModAnthropic(prompt *string) {
	fmt.Printf("\nAnthropic model: %s\n\n", ANTHROPIC_MODEL)
	isSesepuhNeedStream := GetEnv("SESEPUH_NEED_STREAM", "0")

	if isSesepuhNeedStream == "1" {
		ModAnthropicStream(prompt)
	} else {
		ModAnthropicSync(prompt)
	}
}

func ModAnthropicSync(prompt *string) {
	message, err := clientAnthropic.Messages.New(context.TODO(), anthropic.MessageNewParams{MaxTokens: 1024,
		Messages: []anthropic.MessageParam{{
			Role: anthropic.MessageParamRoleUser,
			Content: []anthropic.ContentBlockParamUnion{{
				OfRequestTextBlock: &anthropic.TextBlockParam{Text: *prompt},
			}},
		}},
		Model: ANTHROPIC_MODEL,
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", message.Content)
}

func ModAnthropicStream(prompt *string) {
	stream := clientAnthropic.Messages.NewStreaming(context.TODO(), anthropic.MessageNewParams{
		Model: ANTHROPIC_MODEL,
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(*prompt)),
		},
	})

	message := anthropic.Message{}

	for stream.Next() {
		event := stream.Current()
		err := message.Accumulate(event)
		if err != nil {
			panic(err)
		}

		switch eventVariant := event.AsAny().(type) {
		case anthropic.ContentBlockDeltaEvent:
			switch deltaVariant := eventVariant.Delta.AsAny().(type) {
			case anthropic.TextDelta:
				print(deltaVariant.Text)
			}

		}

		if stream.Err() != nil {
			panic(stream.Err())
		}
	}

	fmt.Println() // add newline after the stream ends
}
