package main

import (
	"log/slog"
	"os"

	"github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/app"
	"github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/configuration"
)

const (
	envDev        = "dev"
	envProduction = "prod"
)

func main() {
	cfg := configuration.LoadConfig()

	logger := initLogger(cfg)

	app := app.NewApp(cfg, logger)
	app.Run()

	// TODO: graceful shutdown
}

// initLogger panics if can't initialize
func initLogger(config *configuration.Config) *slog.Logger {
	var handler slog.Handler

	level := slog.Level(0)
	err := level.UnmarshalText([]byte(config.LogLevel))
	if err != nil {
		panic(err.Error())
	}

	switch config.Env {
	case string(envDev):
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	case string(envProduction):
		panic("not implemented")
	}

	log := slog.New(handler)
	log.Debug("Debug level is on.")
	log.Info("Info level is on.")
	log.Warn("Warn level is on.")
	log.Error("Error level is on.")

	return log
}
