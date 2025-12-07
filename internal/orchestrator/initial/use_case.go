package initial

import (
	authUseCase "github.com/wisaitas/grpc-poc/internal/orchestrator/usecase/auth"
	"google.golang.org/grpc"
)

type useCase struct {
	authUseCase *authUseCase.AuthUseCase
}

func newUseCase(
	sdk *SDK,
) *useCase {
	return &useCase{
		authUseCase: authUseCase.NewAuthUseCase(sdk.Validatorx),
	}
}

func (u *useCase) Register(s *grpc.Server) {
	u.authUseCase.RegisterUseCase(s)
}
