package ai

import (
	"encoding/json"
	"errors"
	"go-microservice/config"
	"go-microservice/models"

	"github.com/go-resty/resty/v2"
)

func GenerateCommitWithPythonService(request models.CommitRequest, gitData string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"git_data":      gitData,
			"model":         request.Model,
			"commit_length": request.CommitLength,
		}).
		Post(config.PythonServiceURL + "/generate")

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", errors.New("Error from Python service: " + resp.String())
	}

	var result map[string]string
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	commitMessage, ok := result["commit_message"]
	if !ok {
		return "", errors.New("Invalid response from Python service")
	}

	return commitMessage, nil
}
