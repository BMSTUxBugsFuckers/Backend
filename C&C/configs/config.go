package configs

type ServiceConfig struct {
	Addr string `yaml:"addr"`
}

func NewServiceConfig() ServiceConfig {
	return ServiceConfig{
		Addr: "127.0.0.1:8080",
	}
}
