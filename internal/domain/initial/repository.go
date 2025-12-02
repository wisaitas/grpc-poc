package initial

import (
	appRepository "github.com/wisaitas/grpc-poc/internal/domain/repository"
)

type repository struct {
	userRepository        appRepository.UserRepository
	userHistoryRepository appRepository.UserHistoryRepository
}

func newRepository(client *client) *repository {
	return &repository{
		userRepository:        appRepository.NewUserRepository(client.postgres),
		userHistoryRepository: appRepository.NewUserHistoryRepository(client.postgres),
	}
}
