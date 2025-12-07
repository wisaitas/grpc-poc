package orchestrator

var Config struct {
	Service struct {
		Port string `env:"PORT" envDefault:"50052"`
		Name string `env:"NAME" envDefault:"orchestrator"`
	} `envPrefix:"SERVICE_"`
	Otel struct {
		Endpoint string `env:"ENDPOINT" envDefault:"localhost:4317"`
	} `envPrefix:"OTEL_"`
}
