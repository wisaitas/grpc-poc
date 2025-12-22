package initial

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/wisaitas/grpc-poc/internal/orchestrator"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	if err := godotenv.Load("./orchestrator.env"); err != nil {
		log.Println(err)
	}

	if err := env.Parse(&orchestrator.Config); err != nil {
		log.Fatalln(err)
	}
}

type App struct {
	server *grpc.Server
	client *client
}

func New() *App {
	client := newClient()
	sdk := newSDK()
	useCase := newUseCase(sdk, client.domainService)

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	useCase.Register(grpcServer)

	reflection.Register(grpcServer)

	return &App{
		client: client,
		server: grpcServer,
	}
}

func (a *App) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", orchestrator.Config.Service.Port))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s service listening on port %s", orchestrator.Config.Service.Name, orchestrator.Config.Service.Port)
	if err := a.server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) Stop() {
	if a.client.otel != nil {
		if err := a.client.otel.Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}

	if err := a.client.domainService.Close(); err != nil {
		log.Fatalln(err)
	}

	a.server.GracefulStop()

	log.Printf("%s service stopped", orchestrator.Config.Service.Name)
}
