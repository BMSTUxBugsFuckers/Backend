package configs

import "github.com/spf13/viper"

type ServiceConfig struct {
	Addr string `yaml:"addr"`
}

func NewServiceConfig() ServiceConfig {
	return ServiceConfig{
		Addr: "127.0.0.1:8080",
	}
}

func InitConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
