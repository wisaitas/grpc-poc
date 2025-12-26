package initial

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	"github.com/wisaitas/grpc-poc/internal/domain"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	if err := godotenv.Load("./domain.env"); err != nil {
		log.Println(err)
	}

	if err := env.Parse(&domain.Config); err != nil {
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
	repository := newRepository(client)
	useCase := newUseCase(sdk, repository)

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
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", domain.Config.Service.Port))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s service listening on port %s", domain.Config.Service.Name, domain.Config.Service.Port)
	if err := a.server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) Stop() {
	a.server.GracefulStop()

	if a.client.otel != nil {
		if err := a.client.otel.Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}

	if err := postgres.Close(a.client.postgres); err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s service stopped", domain.Config.Service.Name)
}
