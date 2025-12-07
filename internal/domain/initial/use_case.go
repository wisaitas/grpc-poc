package initial

import (
	userUseCase "github.com/wisaitas/grpc-poc/internal/domain/usecase/user"
	"google.golang.org/grpc"
)

type useCase struct {
	userUseCase *userUseCase.UserUseCase
}

func newUseCase(
	sdk *SDK,
	repo *Repository,
) *useCase {
	return &useCase{
		userUseCase: userUseCase.NewUserUseCase(sdk.Validatorx, repo.UserRepository, repo.UserHistoryRepository),
	}
}

func (u *useCase) Register(s *grpc.Server) {
	u.userUseCase.RegisterUseCase(s)
}
