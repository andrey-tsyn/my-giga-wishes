package app

import (
	"github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/configuration"
	"github.com/rs/zerolog"
)

type App struct {
	config *configuration.Config
	logger *zerolog.Logger
}

func NewApp(cfg *configuration.Config, logger *zerolog.Logger) *App {
	return &App{
		config: cfg,
		logger: logger,
	}
}

func (a *App) Run() error {

	return nil
}
