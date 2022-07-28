package queries

import (
	"github.com/google/uuid"
	"time"
)

type NewResponse struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
