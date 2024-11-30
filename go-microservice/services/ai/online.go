package ai

import (
	"context"
	"encoding/json"
	"errors"
	"go-microservice/config"
	"go-microservice/models"

	"github.com/go-resty/resty/v2"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateCommitWithOpenAI(request models.CommitRequest, gitData string) (string, error) {
	apiKey := request.OpenAIKey
	if apiKey == "" {
		apiKey = config.OpenAIKey
	}

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	prompt := createOpenAIPrompt(gitData, request.CommitLength)

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: request.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant that generates commit messages based on git diffs.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("No response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

func createOpenAIPrompt(gitData, commitLength string) string {
	lengthInstruction := "short"
	if commitLength == "long" {
		lengthInstruction = "detailed"
	}
	return "Generate a " + lengthInstruction + " commit message based on the following git diff:\n\n" + gitData
}

func GenerateCommitWithHuggingFaceAPI(request models.CommitRequest, gitData string) (string, error) {
	client := resty.New()

	// Use a API da Hugging Face
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer YOUR_HUGGINGFACE_API_TOKEN").
		SetBody(map[string]interface{}{
			"inputs": gitData,
		}).
		Post("https://api-inference.huggingface.co/models/" + request.Model)

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", errors.New("Error from Hugging Face API: " + resp.String())
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", errors.New("No response from Hugging Face API")
	}

	generatedText, ok := result[0]["generated_text"].(string)
	if !ok {
		return "", errors.New("Invalid response format from Hugging Face API")
	}

	return generatedText, nil
}
