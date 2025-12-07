package initial

import (
	"fmt"
	"log"
	"net"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/wisaitas/grpc-poc/internal/domain" // เปลี่ยนจาก router เป็น handler
	"github.com/wisaitas/grpc-poc/internal/orchestrator"
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
	useCase := newUseCase(sdk)

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

	log.Printf("%s service listening on port %s", orchestrator.Config.Service.Name, orchestrator.Config.Service.Port)
	if err := a.server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) Stop() {
	a.server.GracefulStop()

	log.Printf("%s service stopped", orchestrator.Config.Service.Name)
}
