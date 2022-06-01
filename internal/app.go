package internal

import (
	"log"
)

type App struct {
	Logger  *log.Logger
	Command *CommandsBootstrap
}

func NewApp(
	logger *log.Logger,
	cmd *CommandsBootstrap,
) *App {
	return &App{
		Logger:  logger,
		Command: cmd,
	}
}
