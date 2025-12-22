package initial

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wisaitas/grpc-poc/internal/orchestrator"
	pbgen "github.com/wisaitas/grpc-poc/internal/orchestrator/pb/gen"
	"github.com/wisaitas/grpc-poc/pkg/grpcx"
	"github.com/wisaitas/grpc-poc/pkg/otel"
)

type client struct {
	otel          *otel.Telemetry
	domainService *grpcx.GRPCConn[pbgen.DomainServiceClient]
}

func newClient() *client {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	telemetry, err := otel.NewTelemetryGRPC(ctxTimeout, orchestrator.Config.Service.Name, fmt.Sprintf("%s:%d", orchestrator.Config.Otel.Host, orchestrator.Config.Otel.Port))
	if err != nil {
		log.Fatalln(err)
	}

	domainConn, err := grpcx.NewGRPCConn(
		fmt.Sprintf("%s:%d", orchestrator.Config.DomainService.Host, orchestrator.Config.DomainService.Port),
		pbgen.NewDomainServiceClient,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return &client{
		otel:          telemetry,
		domainService: domainConn,
	}
}
