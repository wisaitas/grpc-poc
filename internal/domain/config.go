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
		Host string `env:"HOST" envDefault:"localhost"`
		Port int    `env:"PORT" envDefault:"4317"`
	} `envPrefix:"OTEL_"`
}
