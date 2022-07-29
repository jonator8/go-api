package internal

import (
	"github.com/jonator8/go-api/internal/queries"
	"github.com/jonator8/go-api/internal/repositories"
	"github.com/uptrace/bun"
)

type QueriesBootstrap struct {
	GetNews queries.GetNewsQueryHandler
}

func NewQueriesBootstrap(db *bun.DB) *QueriesBootstrap {
	readNewsRepository := repositories.NewReadNewsRepository(db)

	return &QueriesBootstrap{
		GetNews: queries.NewGetNewsQueryHandler(readNewsRepository),
	}
}
