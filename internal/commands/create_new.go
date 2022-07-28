package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/jonator8/go-api/internal/models"
	"github.com/jonator8/go-api/internal/queries"
	"github.com/jonator8/go-api/internal/repositories"
	"time"
)

type CreateNewCommand struct {
	Title string
	Body  string
}

type CreateNewCommandHandler interface {
	Handle(ctx context.Context, cmd CreateNewCommand) (queries.NewResponse, error)
}

type CreateNewCommandHandlerImpl struct {
	wr repositories.WriteNewsRepository
}

func NewCreateNewCommandHandler(
	wr repositories.WriteNewsRepository,
) CreateNewCommandHandler {
	return &CreateNewCommandHandlerImpl{
		wr: wr,
	}
}

func (h CreateNewCommandHandlerImpl) Handle(ctx context.Context, cmd CreateNewCommand) (queries.NewResponse, error) {
	id := uuid.New()
	now := time.Now()
	err := h.wr.Save(ctx, models.New{
		Id:        id,
		Title:     cmd.Title,
		Body:      cmd.Body,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return queries.NewResponse{}, err
	}

	return queries.NewResponse{
		Id:        id,
		Title:     cmd.Title,
		Body:      cmd.Body,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
