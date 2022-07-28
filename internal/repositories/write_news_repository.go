package repositories

import (
	"context"
	"github.com/jonator8/go-api/internal/models"
	"github.com/uptrace/bun"
)

type WriteNewsRepository interface {
	Seed(ctx context.Context) error
	Save(ctx context.Context, new models.New) error
}

type PostgresWriteNewsRepository struct {
	Db *bun.DB
}

func NewWriteNewsRepository(db *bun.DB) WriteNewsRepository {
	return &PostgresWriteNewsRepository{Db: db}
}

func (r *PostgresWriteNewsRepository) Seed(ctx context.Context) error {
	_, err := r.Db.NewCreateTable().Model((*models.New)(nil)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresWriteNewsRepository) Save(ctx context.Context, new models.New) error {
	_, err := r.Db.NewInsert().Model(&new).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
