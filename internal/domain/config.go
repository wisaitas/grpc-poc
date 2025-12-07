package domain

import "github.com/wisaitas/grpc-poc/pkg/db/postgres"

var Config struct {
	Service struct {
		Port string `env:"PORT" envDefault:"50051"`
		Name string `env:"NAME" envDefault:"domain"`
	} `envPrefix:"SERVICE_"`
	Postgres struct {
		postgres.Config
	} `envPrefix:"POSTGRES_"`
	Otel struct {
		Endpoint string `env:"ENDPOINT" envDefault:"localhost:4317"`
	} `envPrefix:"OTEL_"`
}
