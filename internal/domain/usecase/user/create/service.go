package create

import (
	"context"
	"log/slog"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/domain/repository"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres/entity"
	"github.com/wisaitas/grpc-poc/pkg/otel"
)

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*pb.CreateUserResponse, error)
}

type service struct {
	userRepo        repository.UserRepository
	userHistoryRepo repository.UserHistoryRepository
	logger          *otel.Logger
}

func NewService(
	userRepo repository.UserRepository,
	userHistoryRepo repository.UserHistoryRepository,
	logger *otel.Logger,
) Service {
	return &service{
		userRepo:        userRepo,
		userHistoryRepo: userHistoryRepo,
		logger:          logger,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.logger.Info(ctx, "create user started",
		slog.String("email", req.Email),
		slog.String("firstname", req.FirstName),
	)

	user := mapRequestToEntity(req)

	if err := s.userRepo.Transaction(ctx, func(ctxTx context.Context) error {
		if err := s.userRepo.Create(ctxTx, user); err != nil {
			s.logger.Error(ctx, "failed to create user", slog.String("error", err.Error()))
			return err
		}

		s.logger.Info(ctx, "user created in database", slog.String("user_id", user.ID.String()))

		history := &entity.UserHistory{
			UserID: user.ID,
			Action: UserCreatedAction,
		}

		if err := s.userHistoryRepo.Create(ctxTx, history); err != nil {
			s.logger.Error(ctx, "failed to create user history", slog.String("error", err.Error()))
			return err
		}

		s.logger.Info(ctx, "user history created", slog.String("user_id", user.ID.String()))

		return nil
	}); err != nil {
		s.logger.Error(ctx, "create user failed", slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info(ctx, "create user completed successfully", slog.String("user_id", user.ID.String()))

	return &pb.CreateUserResponse{
		Id: user.ID.String(),
	}, nil
}
