package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type New struct {
	bun.BaseModel `bun:"table:news,alias:n"`

	Id        uuid.UUID `bun:"id,pk,type:uuid"`
	Title     string
	Body      string
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
