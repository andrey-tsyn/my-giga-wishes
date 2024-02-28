package configuration

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env:"ENVIRONMENT" env-required:"true"`
	LogLevel   string `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`
	GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env:"SERVER_GRPC_PORT" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env:"SERVER_TIMEOUT" env-required:"true"`
}

// LoadConfig returns Config, parsed from yaml config file with path in CONFIG_PATH environment variable.
// Panics if cant get config
func LoadConfig() *Config {
	configPath := getConfigPath()

	if _, err := os.Stat(configPath); err != nil {
		panic("can't read configuration file: " + err.Error())
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("can't parse configuration file: " + err.Error())
	}

	return &cfg
}

func getConfigPath() string {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("config path is empty. set CONFIG_PATH environment variable.")
	}
	return configPath
}
