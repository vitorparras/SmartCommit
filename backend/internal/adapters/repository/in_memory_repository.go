package repository

import (
	"backend/internal/core/domain"
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
