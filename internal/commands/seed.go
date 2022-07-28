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
	Handle(ctx context.Context) error
}

type SeedCommandHandlerImpl struct {
	writeNewsRepository repositories.WriteNewsRepository
}

func NewSeedCommandHandler(
	wrNews repositories.WriteNewsRepository,
) SeedCommandHandler {
	return &SeedCommandHandlerImpl{
		writeNewsRepository: wrNews,
	}
}

func (h *SeedCommandHandlerImpl) Handle(ctx context.Context) error {
	err := h.writeNewsRepository.Seed(ctx)
	if err != nil {
		return err
	}

	return nil
}
