package initial

import (
	appRepository "github.com/wisaitas/grpc-poc/internal/domain/repository"
)

type Repository struct {
	UserRepository        appRepository.UserRepository
	UserHistoryRepository appRepository.UserHistoryRepository
}

func newRepository(client *client) *Repository {
	return &Repository{
		UserRepository:        appRepository.NewUserRepository(client.postgres),
		UserHistoryRepository: appRepository.NewUserHistoryRepository(client.postgres),
	}
}
