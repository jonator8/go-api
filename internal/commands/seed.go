package commands

import (
	"context"
	"github.com/jonator8/go-api/internal/repositories"
)

type SeedCommand struct {
	News     bool
	Comments bool
}

type SeedCommandHandler interface {
	Handle(ctx context.Context, cmd SeedCommand) error
}

type SeedCommandHandlerImpl struct {
	writeNewsRepository repositories.WriteTeamRepository
}

func NewSeedCommandHandler(
	wrNews repositories.WriteTeamRepository,
) SeedCommandHandler {
	return &SeedCommandHandlerImpl{
		writeNewsRepository: wrNews,
	}
}

func (h *SeedCommandHandlerImpl) Handle(ctx context.Context, cmd SeedCommand) error {
	return nil
}