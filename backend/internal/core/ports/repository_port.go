package ports

import "backend/internal/core/domain"

type RepositoryPort interface {
	SaveCommit(commit domain.Commit) error
	GetCommits() ([]domain.Commit, error)
}
