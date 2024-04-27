package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PgxDB struct {
	Pool *pgxpool.Pool
}

func (p *PgxDB) GetPool() *pgxpool.Pool {
	return p.Pool
}

func (p *PgxDB) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.Pool.Ping(ctx)
}

func (p *PgxDB) Disconnect() {
	p.Pool.Close()
}

func NewPgxPool(connectionString string) (*PgxDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}
	return &PgxDB{
		Pool: pool,
	}, nil
}
