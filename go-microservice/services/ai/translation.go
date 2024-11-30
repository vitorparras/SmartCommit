package ai

import (
	"encoding/json"
	"errors"
	"go-microservice/config"

	"github.com/go-resty/resty/v2"
)

func TranslateMessage(message, targetLang string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"message":     message,
			"target_lang": targetLang,
		}).
		Post(config.PythonServiceURL + "/translate")

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", errors.New("Error from translation service: " + resp.String())
	}

	var result map[string]string
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	translatedMessage, ok := result["translated_message"]
	if !ok {
		return "", errors.New("Invalid response from translation service")
	}

	return translatedMessage, nil
}
