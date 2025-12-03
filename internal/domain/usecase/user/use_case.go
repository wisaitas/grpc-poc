package user

import (
	"context"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/domain/repository"
	"github.com/wisaitas/grpc-poc/internal/domain/usecase/user/create"
	"github.com/wisaitas/grpc-poc/internal/domain/usecase/user/getlist"
	"github.com/wisaitas/grpc-poc/pkg/validatorx"
	"google.golang.org/grpc"
)

type UserUseCase struct {
	pb.UnimplementedDomainServiceServer
	userCreateHandler  *create.Handler
	userGetListHandler *getlist.Handler
}

func NewUserUseCase(
	validatorx validatorx.Validator,
	userRepo repository.UserRepository,
	userHistoryRepo repository.UserHistoryRepository,
) *UserUseCase {
	return &UserUseCase{
		userCreateHandler:  create.NewHandler(create.NewService(userRepo, userHistoryRepo), validatorx),
		userGetListHandler: getlist.NewHandler(getlist.NewService(userRepo)),
	}
}

func (u *UserUseCase) Register(s *grpc.Server) {
	pb.RegisterDomainServiceServer(s, u)
}

func (u *UserUseCase) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return u.userCreateHandler.CreateUser(ctx, req)
}

func (u *UserUseCase) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	return u.userGetListHandler.GetUserList(ctx, req)
}
