package domain

import "github.com/wisaitas/grpc-poc/pkg/db/postgres"

var Config struct {
	Service struct {
		Port string `env:"PORT" envDefault:"50051"`
	} `envPrefix:"SERVER_"`
	Postgres struct {
		postgres.Config
	} `envPrefix:"POSTGRES_"`
}
