package create

import (
	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres/entity"
)

func mapProtoToRequest(req *pb.CreateUserRequest) *CreateUserRequest {
	return &CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}

func mapRequestToEntity(req *CreateUserRequest) *entity.User {
	return &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}
