package models

import "github.com/google/uuid"

type New struct {
	Id       uuid.UUID
	Title    string
	Body     string
	comments []Comment
}
