package handler

import (
	"context"
	"log"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	userCreate "github.com/wisaitas/grpc-poc/internal/domain/usecase/user/create"
	userGetList "github.com/wisaitas/grpc-poc/internal/domain/usecase/user/getlist"
	"google.golang.org/grpc"
)

// Handler implements DomainServiceServer
type Handler struct {
	pb.UnimplementedDomainServiceServer
	userCreateHandler  *userCreate.Handler
	userGetListHandler *userGetList.Handler
}

// NewHandler creates a new Handler with dependencies injected
func NewHandler(
	userCreateHandler *userCreate.Handler,
	userGetListHandler *userGetList.Handler,
) *Handler {
	return &Handler{
		userCreateHandler:  userCreateHandler,
		userGetListHandler: userGetListHandler,
	}
}

// Register registers the Handler to the gRPC server
func (h *Handler) Register(s *grpc.Server) {
	pb.RegisterDomainServiceServer(s, h)
}

// --- gRPC Implementation ---

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("[Handler] CreateUser called")
	return h.userCreateHandler.CreateUser(ctx, req)
}

func (h *Handler) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	log.Println("[Handler] GetUserList called")
	return h.userGetListHandler.GetUserList(ctx, req)
}
