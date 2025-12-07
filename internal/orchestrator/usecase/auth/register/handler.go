package register

import (
	"context"

	pb "github.com/wisaitas/grpc-poc/internal/orchestrator/pb/gen"
	"github.com/wisaitas/grpc-poc/pkg/validatorx"
)

type Handler struct {
	service    Service
	validatorx validatorx.Validator
}

func NewHandler(
	service Service,
	validatorx validatorx.Validator,
) *Handler {
	return &Handler{
		service:    service,
		validatorx: validatorx,
	}
}

func (h *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}
