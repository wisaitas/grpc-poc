package initial

import (
	pbgen "github.com/wisaitas/grpc-poc/internal/orchestrator/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/orchestrator/usecase/auth"
	authUseCase "github.com/wisaitas/grpc-poc/internal/orchestrator/usecase/auth"
	"github.com/wisaitas/grpc-poc/pkg/grpcx"
	"google.golang.org/grpc"
)

type useCase struct {
	authUseCase *authUseCase.AuthUseCase
}

func newUseCase(
	sdk *SDK,
	domainService *grpcx.GRPCConn[pbgen.DomainServiceClient],
) *useCase {
	return &useCase{
		authUseCase: auth.NewAuthUseCase(sdk.Validatorx, domainService.Client),
	}
}

func (u *useCase) Register(s *grpc.Server) {
	u.authUseCase.RegisterUseCase(s)
}
