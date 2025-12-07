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
) *useCase {
	return &useCase{
		userUseCase: userUseCase.NewUserUseCase(sdk.Validatorx),
	}
}

func (u *useCase) Register(s *grpc.Server) {
	u.userUseCase.Register(s)
}
