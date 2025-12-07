package create

import (
	"context"
	"log/slog"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/domain/repository"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres/entity"
	"github.com/wisaitas/grpc-poc/pkg/telemetry"
)

var logger = telemetry.NewLogger("user-create-service")

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*pb.CreateUserResponse, error)
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

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*pb.CreateUserResponse, error) {
	logger.Info(ctx, "CreateUser started",
		slog.String("email", req.Email),
		slog.String("firstname", req.FirstName),
	)

	user := mapRequestToEntity(req)

	if err := s.userRepo.Transaction(ctx, func(ctxTx context.Context) error {
		if err := s.userRepo.Create(ctxTx, user); err != nil {
			logger.Error(ctx, "Failed to create user", slog.String("error", err.Error()))
			return err
		}

		logger.Info(ctx, "User created in database", slog.String("user_id", user.ID.String()))

		history := &entity.UserHistory{
			UserID: user.ID,
			Action: UserCreatedAction,
		}

		if err := s.userHistoryRepo.Create(ctxTx, history); err != nil {
			logger.Error(ctx, "Failed to create user history", slog.String("error", err.Error()))
			return err
		}

		logger.Info(ctx, "User history created", slog.String("user_id", user.ID.String()))

		return nil
	}); err != nil {
		logger.Error(ctx, "CreateUser failed", slog.String("error", err.Error()))
		return nil, err
	}

	logger.Info(ctx, "CreateUser completed successfully", slog.String("user_id", user.ID.String()))

	return &pb.CreateUserResponse{
		Id: user.ID.String(),
	}, nil
}
