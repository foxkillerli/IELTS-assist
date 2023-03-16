package openai

import (
	"context"
	"fmt"
	"github.com/foxkillerli/IELTS-assist/config"
	"github.com/sashabaranov/go-openai"
	"strconv"
)

func ArticleEditThroughChat(article string, band int) string {
	c := openai.NewClient(config.OPENAI_TOKEN)
	ctx := context.Background()
	resp, err := c.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "You are an IELTS test studying assistant, " +
						"the user would upload an article, " +
						"and you should re-write the article " +
						"with IELTS' band-" + strconv.Itoa(band) + " standard in about 300 words.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: article,
				},
			},
			MaxTokens: 400,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}
	return resp.Choices[0].Message.Content
}

func ArticleEditSuggestion(article string, band int) string {
	c := openai.NewClient(config.OPENAI_TOKEN)
	ctx := context.Background()
	resp, err := c.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "You are an IELTS test studying assistant, " +
						"the user would upload an article, " +
						"and you should rate the article and make study suggestions" +
						"with IELTS' band-" + strconv.Itoa(band) + " standard in Mandarin.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: article,
				},
			},
			MaxTokens: 800,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}
	return resp.Choices[0].Message.Content
}
