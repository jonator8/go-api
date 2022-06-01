package internal

import "log"

type App struct {
	Logger *log.Logger
}

func NewApp(
	logger *log.Logger,
) *App {
	return &App{
		Logger: logger,
	}
}