package create

import (
	"context"
	"log/slog"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/pkg/otel"
	"github.com/wisaitas/grpc-poc/pkg/validatorx"
)

type Handler struct {
	service    Service
	validatorx validatorx.Validator
	logger     *otel.Logger
}

func NewHandler(
	service Service,
	validatorx validatorx.Validator,
	logger *otel.Logger,
) *Handler {
	return &Handler{
		service:    service,
		validatorx: validatorx,
		logger:     logger,
	}
}

func (h *Handler) CreateUser(ctx context.Context, proto *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	req, err := h.validateRequest(proto)
	if err != nil {
		h.logger.Warn(ctx, "validation failed", slog.String("error", err.Error()))
		return nil, err
	}

	return h.service.CreateUser(ctx, req)
}

func (h *Handler) validateRequest(proto *pb.CreateUserRequest) (*CreateUserRequest, error) {
	createUserRequest := mapProtoToRequest(proto)

	if err := h.validatorx.ValidateStruct(createUserRequest); err != nil {
		return nil, err
	}

	return createUserRequest, nil
}
