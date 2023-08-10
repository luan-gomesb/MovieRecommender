package openaiapi

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)


func Openaiapi(openaiKey string) (string,error) {
	client := openai.NewClient(openaiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			MaxTokens: 500,
			Temperature: 0,
			TopP: 1,
			PresencePenalty: 0.0,	
			FrequencyPenalty: 0.0,
			Messages: []openai.ChatCompletionMessage{openai.ChatCompletionMessage{Content: 
			 "me indique 5 filmes parecidos com esses matrix, jogador numero 1 e prometheus. Retorne essas instruções como um objeto JSON com a estrutura {“Headline”:[“Title”:“string”, “Description”:“string”,“YouTubeTrailler”:“string”,“Date”:“date”]}. Não retorne nenhum texto ou numeração que não seja json.",
			Role: "user",
			}},
			
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

