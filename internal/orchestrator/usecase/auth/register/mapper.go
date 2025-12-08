package register

import pb "github.com/wisaitas/grpc-poc/internal/orchestrator/pb/gen"

func mapProtoToRequest(req *pb.RegisterRequest) *RegisterRequest {
	return &RegisterRequest{
		Email:           req.Email,
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}
}
