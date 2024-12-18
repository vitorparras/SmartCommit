package repository

import (
	"backend/internal/core/domain"
	"backend/internal/core/ports"
	"mime/multipart"
	"os/exec"
	"sync"
)

type InMemoryRepository struct {
	commits []domain.Commit
	mutex   sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		commits: []domain.Commit{},
	}
}

func (r *InMemoryRepository) SaveCommit(commit domain.Commit) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.commits = append(r.commits, commit)
	return nil
}

func (r *InMemoryRepository) GetCommits() ([]domain.Commit, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.commits, nil
}

type CommitService struct {
	ai   ports.AIPort
	repo ports.RepositoryPort
}

// Método para validar a existência do repositório Git
func (s *CommitService) ValidateGitRepository(readmeFile *multipart.FileHeader) (bool, error) {
	// Verifica se o arquivo README existe
	if readmeFile == nil {
		return false, nil
	}

	// Verifica se o arquivo README está em um repositório Git
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	if err != nil {
		return false, err // Não é um repositório Git
	}

	return true, nil // É um repositório Git
}
