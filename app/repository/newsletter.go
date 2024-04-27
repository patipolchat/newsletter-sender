package repository

import (
	"context"
	"energy-response-assignment/config"
	"energy-response-assignment/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Newsletter interface {
	CreateNewsLetter(ctx context.Context, arg sqlc.CreateNewsLetterParams) (sqlc.Newsletter, error)
	GetActiveSubscribersWithPaginate(ctx context.Context, arg sqlc.GetActiveSubscribersWithPaginateParams) ([]sqlc.Subscriber, error)
	CountActiveSubscriber(ctx context.Context) (int64, error)
	CreateSendingEmails(ctx context.Context, arg []sqlc.CreateSendingEmailsParams) *sqlc.CreateSendingEmailsBatchResults
	BatchUpdateSendingEmailsStatus(ctx context.Context, arg []sqlc.BatchUpdateSendingEmailsStatusParams) *sqlc.BatchUpdateSendingEmailsStatusBatchResults
	UpdateSendingEmailsStatus(ctx context.Context, arg sqlc.UpdateSendingEmailsStatusParams) (sqlc.SendingEmail, error)
}

type newletterImpl struct {
	cfg *config.Config
	*sqlc.Queries
}

func NewNewsletter(cfg *config.Config, pool *pgxpool.Pool) Newsletter {
	queries := sqlc.New(pool)
	return &newletterImpl{
		cfg:     cfg,
		Queries: queries,
	}
}
