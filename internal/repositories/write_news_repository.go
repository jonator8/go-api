package repositories

import (
	"context"
	"github.com/uptrace/bun"
)

type WriteTeamRepository interface {
	Seed(ctx context.Context) error
}

type PostgresWriteNewsRepository struct {
	Db *bun.DB
}

func NewWriteNewsRepository(db *bun.DB) WriteTeamRepository {
	return &PostgresWriteNewsRepository{Db: db}
}

func (r *PostgresWriteNewsRepository) Seed(ctx context.Context) error {
	return nil
}
