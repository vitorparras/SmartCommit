package ai

import (
	"errors"
	"go-microservice/config"
	"go-microservice/models"
)

var onlineModels = []string{"gpt-3.5-turbo", "gpt-4", "EleutherAI/gpt-j-6B", "EleutherAI/gpt-neo-2.7B", "facebook/opt-125m"}
var offlineModels = []string{"google/flan-t5-base", "bigscience/bloom-560m"}

func GetAvailableModels() map[string][]string {
	return map[string][]string{
		"online":  onlineModels,
		"offline": offlineModels,
	}
}

func IsOnlineModel(model string) bool {
	for _, m := range onlineModels {
		if m == model {
			return true
		}
	}
	return false
}

func GenerateCommitMessageOnline(request models.CommitRequest, gitData string) (string, error) {
	if request.OpenAIKey == "" && config.OpenAIKey == "" {
		return "", errors.New("OpenAI API key is required for online models")
	}
	// Se o modelo for da OpenAI
	if request.Model == "gpt-3.5-turbo" || request.Model == "gpt-4" {
		return GenerateCommitWithOpenAI(request, gitData)
	} else {
		// Caso contrário, use modelos da Hugging Face
		return GenerateCommitWithHuggingFaceAPI(request, gitData)
	}
}

func GenerateCommitMessageOffline(request models.CommitRequest, gitData string) (string, error) {
	// Chama o microserviço Python
	return GenerateCommitWithPythonService(request, gitData)
}
