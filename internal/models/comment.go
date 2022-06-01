package models

import "github.com/google/uuid"

type Comment struct {
	id      uuid.UUID
	user    string
	message string
}
