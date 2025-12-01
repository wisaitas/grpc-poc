package initial

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/wisaitas/grpc-poc/internal/domain"
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
	Server *grpc.Server
	client *client
}

func New() *App {
	client := newClient()
	return &App{
		client: client,
	}
}
