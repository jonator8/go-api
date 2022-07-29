package queries

import (
	"context"
	"github.com/google/uuid"
	"github.com/jonator8/go-api/internal/repositories"
	"time"
)

type NewResponse struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetNewsQueryHandler interface {
	Handle(ctx context.Context) ([]NewResponse, error)
}

type GetNewsQueryHandlerImpl struct {
	rr repositories.ReadNewsRepository
}

func NewGetNewsQueryHandler(
	rr repositories.ReadNewsRepository,
) GetNewsQueryHandler {
	return &GetNewsQueryHandlerImpl{
		rr: rr,
	}
}

func (h GetNewsQueryHandlerImpl) Handle(ctx context.Context) ([]NewResponse, error) {
	news := make([]NewResponse, 0)
	newsModel, err := h.rr.FindAll(ctx)
	if err != nil {
		return news, err
	}

	for _, n := range newsModel {
		news = append(news, NewResponse{
			Id:        n.Id,
			Title:     n.Title,
			Body:      n.Body,
			CreatedAt: n.CreatedAt,
			UpdatedAt: n.UpdatedAt,
		})
	}

	return news, nil
}