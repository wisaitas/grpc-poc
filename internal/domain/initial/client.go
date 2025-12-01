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
	postgres, err := postgres.NewPostgreSQL(postgres.Config{
		Host:            domain.Config.Postgres.Host,
		Port:            domain.Config.Postgres.Port,
		User:            domain.Config.Postgres.User,
		Password:        domain.Config.Postgres.Password,
		DBName:          domain.Config.Postgres.DBName,
		SSLMode:         domain.Config.Postgres.SSLMode,
		MaxIdleConns:    domain.Config.Postgres.MaxIdleConns,
		MaxOpenConns:    domain.Config.Postgres.MaxOpenConns,
		ConnMaxLifetime: domain.Config.Postgres.ConnMaxLifetime,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &client{
		postgres: postgres,
	}
}
