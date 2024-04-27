package repository

import (
	"context"
	"energy-response-assignment/config"
	"energy-response-assignment/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Subscriber interface {
	UpsertSubscriber(ctx context.Context, arg sqlc.UpsertSubscriberParams) (sqlc.Subscriber, error)
}

type subscriberImpl struct {
	cfg *config.Config
	*sqlc.Queries
}

func NewSubscriber(cfg *config.Config, pool *pgxpool.Pool) Subscriber {
	queries := sqlc.New(pool)
	return &subscriberImpl{
		cfg:     cfg,
		Queries: queries,
	}
}
