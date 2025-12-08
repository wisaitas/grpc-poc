package initial

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/wisaitas/grpc-poc/internal/domain/pb/gen"
	"github.com/wisaitas/grpc-poc/internal/orchestrator"
	"github.com/wisaitas/grpc-poc/pkg/otel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	otel         *otel.Telemetry
	domainClient pb.DomainServiceClient
	domainConn   *grpc.ClientConn
}

func newClient() *client {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	telemetry, err := otel.NewTelemetryGRPC(ctxTimeout, orchestrator.Config.Service.Name, fmt.Sprintf("%s:%d", orchestrator.Config.Otel.Host, orchestrator.Config.Otel.Port))
	if err != nil {
		log.Fatalln(err)
	}

	domainConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", orchestrator.Config.DomainService.Host, orchestrator.Config.DomainService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()), // เพิ่ม OTel tracing
	)
	if err != nil {
		log.Fatalln("failed to connect to domain service:", err)
	}

	return &client{
		otel:         telemetry,
		domainClient: pb.NewDomainServiceClient(domainConn),
		domainConn:   domainConn,
	}
}
