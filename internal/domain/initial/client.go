package initial

import (
	"log"

	"github.com/wisaitas/grpc-poc/internal/domain"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
	"gorm.io/gorm"
)

type client struct {
	postgres *gorm.DB
}

func newClient() *client {
	postgres, err := postgres.NewPostgreSQL(domain.Config.Postgres.Config)
	if err != nil {
		log.Fatalln(err)
	}

	return &client{
		postgres: postgres,
	}
}
