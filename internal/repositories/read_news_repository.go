package repositories

import (
	"context"
	"github.com/jonator8/go-api/internal/models"
	"github.com/uptrace/bun"
)

type ReadNewsRepository interface {
	FindAll(ctx context.Context) ([]models.New, error)
}

type PostgresReadNewsRepository struct {
	Db *bun.DB
}

func NewReadNewsRepository(db *bun.DB) ReadNewsRepository {
	return &PostgresReadNewsRepository{Db: db}
}

func (r *PostgresReadNewsRepository) FindAll(ctx context.Context) ([]models.New, error) {
	var news []models.New
	err := r.Db.NewSelect().Model(&news).Scan(ctx)
	if err != nil {
		return news, err
	}

	return news, nil
}
