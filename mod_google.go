package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var GEMINI_API_KEY = os.Getenv("GEMINI_API_KEY")
var __GOOGLE_MODEL = GetModelToUse("__GOOGLE_MODEL", "gemini-2.0-flash")

func ModGoogle(prompt *string) {
	if os.Getenv("SESEPUH_HUB_RES_ONLY") != "1" {
		fmt.Printf("\nGoogle model: %s\n\n", __GOOGLE_MODEL)
	}

	isSesepuhNeedStream := GetEnv("SESEPUH_NEED_STREAM", "0")

	if isSesepuhNeedStream == "1" {
		ModGoogleStream(prompt)
	} else {
		ModGoogleSync(prompt)
	}
}

func ModGoogleSync(prompt *string) {
	// fmt.Println("[DEBUG] ModGoogleSync", prompt)

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel(__GOOGLE_MODEL)
	resp, err := model.GenerateContent(ctx, genai.Text(*prompt))

	if err != nil {
		log.Fatal(err)
	}

	if resp.Candidates != nil {
		for _, v := range resp.Candidates {
			for _, k := range v.Content.Parts {
				fmt.Println(k.(genai.Text))
			}
		}
	}
}

func ModGoogleStream(prompt *string) {
	// fmt.Println("[DEBUG] ModGoogleStream", prompt)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel(__GOOGLE_MODEL)

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

		// print resp to human-readable format
		if resp.Candidates != nil {
			for _, v := range resp.Candidates {
				for _, k := range v.Content.Parts {
					fmt.Print(k.(genai.Text))
				}
			}
		}
	}
}
