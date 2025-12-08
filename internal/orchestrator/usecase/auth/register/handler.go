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

func (h *Handler) Register(ctx context.Context, proto *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	req, err := h.validateRequest(proto)
	if err != nil {
		return nil, err
	}

	return h.service.Register(ctx, req)
}

func (h *Handler) validateRequest(proto *pb.RegisterRequest) (*RegisterRequest, error) {
	req := mapProtoToRequest(proto)

	if err := h.validatorx.ValidateStruct(req); err != nil {
		return nil, err
	}

	return req, nil
}
