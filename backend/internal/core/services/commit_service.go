package services

import (
	"backend/internal/core/domain"
	"backend/internal/core/ports"
	"io"
	"mime/multipart"
	"time"
)

type CommitService struct {
	ai   ports.AIPort
	repo ports.RepositoryPort
}

func NewCommitService(ai ports.AIPort, repo ports.RepositoryPort) *CommitService {
	return &CommitService{ai: ai, repo: repo}
}

func (s *CommitService) GenerateAndSaveCommit(description string, author string, config domain.CommitConfig) (domain.Commit, error) {
	readmeContent, err := readFileContent(config.ReadmeFile)
	if err != nil {
		return domain.Commit{}, err
	}

	message, err := s.ai.GenerateCommitMessage(description, readmeContent, config)
	if err != nil {
		return domain.Commit{}, err
	}

	commit := domain.Commit{
		Message: message,
		Author:  author,
		Date:    time.Now().Format(time.RFC3339),
		Config:  config,
	}

	err = s.repo.SaveCommit(commit)
	if err != nil {
		return domain.Commit{}, err
	}

	return commit, nil
}

func (s *CommitService) GetAllCommits() ([]domain.Commit, error) {
	return s.repo.GetCommits()
}

func (s *CommitService) GetAvailableAIModels() []domain.Model {
	return s.ai.GetAvailableModels()
}

func readFileContent(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	content, err := io.ReadAll(src)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
