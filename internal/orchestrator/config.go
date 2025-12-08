package orchestrator

var Config struct {
	Service struct {
		Port string `env:"PORT" envDefault:"50052"`
		Name string `env:"NAME" envDefault:"orchestrator"`
	} `envPrefix:"SERVICE_"`
	Otel struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port int    `env:"PORT" envDefault:"4317"`
	} `envPrefix:"OTEL_"`
	DomainService struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port int    `env:"PORT" envDefault:"50051"`
	} `envPrefix:"DOMAIN_SERVICE_"`
}
