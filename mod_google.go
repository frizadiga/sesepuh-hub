package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var GEMINI_API_KEY = os.Getenv("GEMINI_API_KEY")
var GOOGLE_MODEL = GetModelToUse("GOOGLE_MODEL", "gemini-2.0-flash")
var resultBufGoogle strings.Builder

func ModGoogle(prompt *string) {
	if os.Getenv("SESEPUH_HUB_RES_ONLY") != "1" {
		fmt.Printf("\nGoogle model: %s\n\n", GOOGLE_MODEL)
	}

	isStreaming := GetEnv("SESEPUH_HUB_STREAMING", "0")

	if isStreaming == "1" {
		ModGoogleStream(prompt)
	} else {
		ModGoogleSync(prompt)
	}
}

func ModGoogleSync(prompt *string) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel(GOOGLE_MODEL)
	resp, err := model.GenerateContent(ctx, genai.Text(*prompt))

	if err != nil {
		log.Fatal(err)
	}

	if resp.Candidates != nil {
		for _, c := range resp.Candidates {
			for _, part := range c.Content.Parts {
				switch p := part.(type) {
				case genai.Text:
					fmt.Println(p)
					resultBufGoogle.WriteString(string(p))
				default:
					continue // skip non-text parts
				}
			}
		}
	}

	WriteRespToFile([]byte(resultBufGoogle.String()), "")
}

func ModGoogleStream(prompt *string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel(GOOGLE_MODEL)

	finalPrompt := genai.Text(*prompt)
	iter := model.GenerateContentStream(ctx, finalPrompt)

	for {
		resp, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if resp.Candidates != nil {
			for _, c := range resp.Candidates {
				for _, part := range c.Content.Parts {
					switch p := part.(type) {
					case genai.Text:
						fmt.Print(p)
						resultBufGoogle.WriteString(string(p))
					default:
						continue // skip non-text parts
					}
				}
			}
		}
	}

	WriteRespToFile([]byte(resultBufGoogle.String()), "")
}
