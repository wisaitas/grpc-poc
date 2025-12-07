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
	"github.com/wisaitas/grpc-poc/pkg/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// func init() {
// 	if err := godotenv.Load("./domain.env"); err != nil {
// 		log.Println(err)
// 	}

// 	if err := env.Parse(&domain.Config); err != nil {
// 		log.Fatalln(err)
// 	}
// }

// type App struct {
// 	server *grpc.Server
// 	client *client
// }

// func New() *App {
// 	client := newClient()
// 	sdk := newSDK()
// 	repository := newRepository(client)
// 	useCase := newUseCase(sdk, repository)

// 	grpcServer := grpc.NewServer()
// 	useCase.Register(grpcServer)

// 	// for postman to see methods
// 	reflection.Register(grpcServer)

// 	return &App{
// 		client: client,
// 		server: grpcServer,
// 	}
// }

// func (a *App) Start() {
// 	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", domain.Config.Service.Port))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Printf("domain service listening on port %s", domain.Config.Service.Port)
// 	if err := a.server.Serve(listen); err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func (a *App) Stop() {
// 	if err := postgres.Close(a.client.postgres); err != nil {
// 		log.Fatalln(err)
// 	}

// 	a.server.GracefulStop()

//		log.Println("domain service stopped")
//	}
func init() {
	if err := godotenv.Load("./domain.env"); err != nil {
		log.Println(err)
	}

	if err := env.Parse(&domain.Config); err != nil {
		log.Fatalln(err)
	}
}

type App struct {
	server    *grpc.Server
	client    *client
	telemetry *telemetry.Telemetry
}

func New() *App {
	ctx := context.Background()

	// Initialize OpenTelemetry (Traces + Logs)
	tel, err := telemetry.Init(ctx, "domain-service", domain.Config.Otel.Endpoint)
	if err != nil {
		log.Printf("Failed to initialize telemetry: %v", err)
	}

	client := newClient()
	sdk := newSDK()
	repository := newRepository(client)
	useCase := newUseCase(sdk, repository)

	// เพิ่ม OpenTelemetry interceptors
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	useCase.Register(grpcServer)
	reflection.Register(grpcServer)

	return &App{
		client:    client,
		server:    grpcServer,
		telemetry: tel,
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
	// Shutdown telemetry (traces + logs)
	if a.telemetry != nil {
		if err := a.telemetry.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down telemetry: %v", err)
		}
	}

	if err := postgres.Close(a.client.postgres); err != nil {
		log.Fatalln(err)
	}

	a.server.GracefulStop()
	log.Printf("%s service stopped", domain.Config.Service.Name)
}
