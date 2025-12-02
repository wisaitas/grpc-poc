package getlist

import (
	"context"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	return h.service.GetList(ctx, req)
}
