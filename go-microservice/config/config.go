package config

import (
	"os"
)

var (
	ServerPort       string
	OpenAIKey        string
	PythonServiceURL string
)

func LoadConfig() {
	ServerPort = getEnv("SERVER_PORT", ":8080")
	OpenAIKey = os.Getenv("OPENAI_API_KEY")
	PythonServiceURL = getEnv("PYTHON_SERVICE_URL", "http://python-microservice:5000")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
