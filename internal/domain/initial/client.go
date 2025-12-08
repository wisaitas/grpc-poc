package initial

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wisaitas/grpc-poc/internal/domain"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres/entity"
	"github.com/wisaitas/grpc-poc/pkg/otel"
	"gorm.io/gorm"
)

type client struct {
	postgres *gorm.DB
	otel     *otel.Telemetry
}

func newClient() *client {
	postgres, err := postgres.NewPostgreSQL(domain.Config.Postgres.Config)
	if err != nil {
		log.Fatalln(err)
	}

	if err := postgres.AutoMigrate(
		&entity.User{},
		&entity.UserHistory{},
	); err != nil {
		log.Fatalln(err)
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	telemetry, err := otel.NewTelemetryGRPC(ctxTimeout, domain.Config.Service.Name, fmt.Sprintf("%s:%d", domain.Config.Otel.Host, domain.Config.Otel.Port))
	if err != nil {
		log.Fatalln(err)
	}

	return &client{
		postgres: postgres,
		otel:     telemetry,
	}
}
