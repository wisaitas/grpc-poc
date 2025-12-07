package initial

import (
	"fmt"
	"log"
	"net"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/wisaitas/grpc-poc/internal/domain" // เปลี่ยนจาก router เป็น handler
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
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
	useCase := newUseCase(sdk, repository)

	grpcServer := grpc.NewServer()
	useCase.Register(grpcServer)

	// for postman to see methods
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

	log.Printf("domain service listening on port %s", domain.Config.Service.Port)
	if err := a.server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) Stop() {
	if err := postgres.Close(a.client.postgres); err != nil {
		log.Fatalln(err)
	}

	a.server.GracefulStop()

	log.Println("domain service stopped")
}
