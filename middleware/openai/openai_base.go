package openai

import "github.com/sashabaranov/go-openai"

type OpenAIBase struct {
	token  string
	client *openai.Client
}
