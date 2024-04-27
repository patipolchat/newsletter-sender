package db

import (
	"github.com/pashagolub/pgxmock/v3"
	"testing"
)

func NewPgxMock(t *testing.T) pgxmock.PgxPoolIface {
	pool, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error creating pgx mock: %v", err)
	}
	return pool
}
