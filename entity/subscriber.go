package entity

import (
	"github.com/google/uuid"
	"time"
)

type Subscriber struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Active    bool      `db:"active" json:"active"`
}
