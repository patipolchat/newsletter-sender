package entity

import (
	"github.com/google/uuid"
	"time"
)

type NewsLetter struct {
	ID        uuid.UUID
	Subject   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Header    string
	Body      string
}
