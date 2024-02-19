package app

import (
	"log/slog"

	"github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/configuration"
)

type App struct {
	config *configuration.Config
	logger *slog.Logger
}

func NewApp(cfg *configuration.Config, logger *slog.Logger) *App {
	return &App{
		config: cfg,
		logger: logger,
	}
}

func (a *App) Run() error {

	return nil
}
