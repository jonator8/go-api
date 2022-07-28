package internal

import (
	"github.com/jonator8/go-api/internal/commands"
	"github.com/jonator8/go-api/internal/repositories"
	"github.com/uptrace/bun"
)

type CommandsBootstrap struct {
	Seed      commands.SeedCommandHandler
	CreateNew commands.CreateNewCommandHandler
}

func NewCommandsBootstrap(
	db *bun.DB,
) *CommandsBootstrap {
	writeNewsRepository := repositories.NewWriteNewsRepository(db)

	return &CommandsBootstrap{
		Seed: commands.NewSeedCommandHandler(
			writeNewsRepository,
		),
		CreateNew: commands.NewCreateNewCommandHandler(
			writeNewsRepository,
		),
	}
}
