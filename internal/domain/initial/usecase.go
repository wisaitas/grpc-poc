package initial

import (
	userCreate "github.com/wisaitas/grpc-poc/internal/domain/usecase/user/create"
	userGetList "github.com/wisaitas/grpc-poc/internal/domain/usecase/user/getlist"
)

type usecase struct {
	userCreateService  userCreate.Service
	userGetListService userGetList.Service
}

func newUsecase(repo *repository) *usecase {
	return &usecase{
		userCreateService: userCreate.NewService(
			repo.userRepository,
			repo.userHistoryRepository,
		),
		userGetListService: userGetList.NewService(
			repo.userRepository,
		),
	}
}
