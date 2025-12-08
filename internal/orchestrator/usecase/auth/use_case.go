package auth

import (
	"context"

	domainpb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	pb "github.com/wisaitas/grpc-poc/internal/orchestrator/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/orchestrator/usecase/auth/register"
	"github.com/wisaitas/grpc-poc/pkg/otel"
	"github.com/wisaitas/grpc-poc/pkg/validatorx"
	"google.golang.org/grpc"
)

type AuthUseCase struct {
	pb.UnimplementedOrchestratorServiceServer
	registerHandler *register.Handler
}

func NewAuthUseCase(
	validatorx validatorx.Validator,
	domainClient domainpb.DomainServiceClient,
) *AuthUseCase {
	return &AuthUseCase{
		registerHandler: register.NewHandler(
			register.NewService(otel.NewLogger("register-service"), domainClient),
			validatorx,
		),
	}
}

func (u *AuthUseCase) RegisterUseCase(s *grpc.Server) {
	pb.RegisterOrchestratorServiceServer(s, u)
}

func (u *AuthUseCase) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return u.registerHandler.Register(ctx, req)
}
