package create

import (
	"context"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/domain/repository"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres/entity"
)

type Service interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
}

type service struct {
	userRepo        repository.UserRepository
	userHistoryRepo repository.UserHistoryRepository
}

func NewService(
	userRepo repository.UserRepository,
	userHistoryRepo repository.UserHistoryRepository,
) Service {
	return &service{
		userRepo:        userRepo,
		userHistoryRepo: userHistoryRepo,
	}
}

func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	if err := s.userRepo.Transaction(ctx, func(ctxTx context.Context) error {
		if err := s.userRepo.Create(ctxTx, user); err != nil {
			return err
		}

		history := &entity.UserHistory{
			UserID: user.ID,
			Action: "USER_CREATED",
		}

		if err := s.userHistoryRepo.Create(ctxTx, history); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id: user.ID.String(),
	}, nil
}
