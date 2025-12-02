package initial

import (
	"fmt"
	"log"
	"net"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/wisaitas/grpc-poc/internal/domain"
	"github.com/wisaitas/grpc-poc/internal/domain/handler" // เปลี่ยนจาก router เป็น handler
	userCreate "github.com/wisaitas/grpc-poc/internal/domain/usecase/user/create"
	userGetList "github.com/wisaitas/grpc-poc/internal/domain/usecase/user/getlist"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
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
	usecase := newUsecase(repository)

	// 1. เตรียม UseCase Handlers
	createHandler := userCreate.NewHandler(usecase.userCreateService, sdk.validatorx)
	getListHandler := userGetList.NewHandler(usecase.userGetListService)

	// 2. สร้าง gRPC Handler (เดิมคือ Router)
	appHandler := handler.NewHandler(createHandler, getListHandler)

	// 3. สร้าง gRPC Server และ Register
	grpcServer := grpc.NewServer()
	appHandler.Register(grpcServer)

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

	if err := a.server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) Close() {
	if err := postgres.Close(a.client.postgres); err != nil {
		log.Fatalln(err)
	}
}
