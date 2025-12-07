package getlist

import (
	"context"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/domain/repository"
)

type Service interface {
	GetList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error)
}

type service struct {
	userRepo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) GetList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.UserData
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.UserData{
			Id:        u.ID.String(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		})
	}

	return &pb.GetUserListResponse{
		Users: pbUsers,
		Total: int64(len(pbUsers)),
	}, nil
}
