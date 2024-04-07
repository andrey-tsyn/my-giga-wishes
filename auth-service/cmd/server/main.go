package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/app"
	"github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/configuration"
	"github.com/rs/zerolog"
)

const (
	envDev        = "dev"
	envProduction = "prod"
)

func main() {
	cfg := configuration.LoadConfig()

	// logger := initLogger(cfg)
	logger := zerolog.New(zerolog.NewConsoleWriter())

	app := app.NewApp(cfg, &logger)

	err := app.Run()
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	terminate := make(chan os.Signal, 2)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)
	<-terminate
	// TODO: graceful shutdown
}

// panics if can't initialize
func initLogger(config *configuration.Config) {
}
