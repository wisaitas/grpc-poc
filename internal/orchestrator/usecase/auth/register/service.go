package register

import (
	"context"
	"log/slog"

	domainpb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	pb "github.com/wisaitas/grpc-poc/internal/orchestrator/pb/gen"
	"github.com/wisaitas/grpc-poc/pkg/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) (*pb.RegisterResponse, error)
}

type service struct {
	logger       *otel.Logger
	domainClient domainpb.DomainServiceClient
	tracer       trace.Tracer
}

func NewService(
	logger *otel.Logger,
	domainClient domainpb.DomainServiceClient,
) Service {
	return &service{
		logger:       logger,
		domainClient: domainClient,
		tracer:       otel.GetTracerProvider().Tracer("register-service"),
	}
}

func (s *service) Register(ctx context.Context, req *RegisterRequest) (*pb.RegisterResponse, error) {
	// Start a new span for tracing
	ctx, span := s.tracer.Start(ctx, "Register",
		trace.WithAttributes(
			attribute.String("email", req.Email),
			attribute.String("first_name", req.FirstName),
		),
	)
	defer span.End()

	s.logger.Info(ctx, "register started",
		slog.String("email", req.Email),
		slog.String("first_name", req.FirstName),
	)

	// Call domain service to create user
	createUserReq := &domainpb.CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	s.logger.Info(ctx, "calling domain CreateUser")

	resp, err := s.domainClient.CreateUser(ctx, createUserReq)
	if err != nil {
		s.logger.Error(ctx, "failed to create user in domain",
			slog.String("error", err.Error()),
		)
		span.RecordError(err)
		return nil, err
	}

	s.logger.Info(ctx, "register completed successfully",
		slog.String("user_id", resp.Id),
	)

	return &pb.RegisterResponse{
		Token: resp.Id,
	}, nil
}
