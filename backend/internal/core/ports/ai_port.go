package ports

import "backend/internal/core/domain"

type AIPort interface {
	GenerateCommitMessage(description string, readmeContent string, config domain.CommitConfig) (string, error)
	GetAvailableModels() []domain.Model
}
