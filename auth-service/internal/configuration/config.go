package configuration

import "time"

type Config struct {
	Env      string `yaml:"env" env:"ENVIRONMENT" env-required:"true"`
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env:"SERVER_GRPC_PORT" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env:"SERVER_TIMEOUT" env-required:"true"`
}
