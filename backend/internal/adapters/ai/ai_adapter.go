package ai

import (
	"backend/internal/core/domain"
	"fmt"
)

type AIAdapter struct {
	// Configuração do serviço de IA iria aqui
}

func NewAIAdapter() *AIAdapter {
	return &AIAdapter{}
}

func (a *AIAdapter) GenerateCommitMessage(description string, readmeContent string, config domain.CommitConfig) (string, error) {
	// Aqui você implementaria a chamada real para o seu serviço de IA
	// usando os parâmetros de configuração para personalizar a geração
	return fmt.Sprintf("[%s] AI generated commit for: %s (Style: %s, Lang: %s, Readme length: %d)",
		config.Format,
		description,
		config.Style,
		config.Language,
		len(readmeContent),
	), nil
}

func (a *AIAdapter) GetAvailableModels() []domain.Model {
	return []domain.Model{
		{Id: 2, Name: "gpt-3.5-turbo", Type: domain.Online},
		{Id: 3, Name: "gpt-4", Type: domain.Online},
		{Id: 4, Name: "codex", Type: domain.Online},
		{Id: 5, Name: "claude-2", Type: domain.Online},
		{Id: 6, Name: "local-llama", Type: domain.Offline},
	}
}
