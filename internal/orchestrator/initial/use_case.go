package initial

import (
	domainpb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/orchestrator/usecase/auth"
	authUseCase "github.com/wisaitas/grpc-poc/internal/orchestrator/usecase/auth"
	"google.golang.org/grpc"
)

type useCase struct {
	authUseCase *authUseCase.AuthUseCase
}

func newUseCase(sdk *SDK, domainClient domainpb.DomainServiceClient) *useCase {
	return &useCase{
		authUseCase: auth.NewAuthUseCase(sdk.Validatorx, domainClient),
	}
}

func (u *useCase) Register(s *grpc.Server) {
	u.authUseCase.RegisterUseCase(s)
}
