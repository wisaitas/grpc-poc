package create

import (
	"context"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
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

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := h.validateRequest(req); err != nil {
		return nil, err
	}

	return h.service.CreateUser(ctx, req)
}

func (h *Handler) validateRequest(req *pb.CreateUserRequest) error {
	return h.validatorx.ValidateStruct(req)
}
