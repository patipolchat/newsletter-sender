// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: newsletter.sql

package sqlc

import (
	"context"
)

const createNewsLetter = `-- name: CreateNewsLetter :one
INSERT INTO newsletters (header, body)
VALUES ($1, $2)
RETURNING id, header, body, created_at, updated_at
`

type CreateNewsLetterParams struct {
	Header string
	Body   string
}

// description: Create a new newsletter
// input: header: String, body: String
// output: Newsletter: Object
// example: CreateNewsLetter("Newsletter 1", "This is the first newsletter")
func (q *Queries) CreateNewsLetter(ctx context.Context, arg CreateNewsLetterParams) (Newsletter, error) {
	row := q.db.QueryRow(ctx, createNewsLetter, arg.Header, arg.Body)
	var i Newsletter
	err := row.Scan(
		&i.ID,
		&i.Header,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}