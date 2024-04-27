package sqlc

import (
	"context"
	"energy-response-assignment/util/db"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueries_CountActiveSubscriber(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				db: db.NewPgxMock(t),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    5,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.fields.db.(pgxmock.PgxPoolIface)
			rows := pgxmock.NewRows([]string{"count"}).AddRow(tt.want)
			mock.ExpectQuery("SELECT COUNT").WillReturnRows(rows)
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.CountActiveSubscriber(tt.args.ctx)
			if tt.wantErr(t, err, "CountActiveSubscriber() error = %v, wantErr %v", err, tt.wantErr) {
				assert.Equal(t, tt.want, got, "CountActiveSubscriber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_UpsertSubscriber(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg UpsertSubscriberParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Subscriber
		setupFn func(mock pgxmock.PgxPoolIface, params UpsertSubscriberParams, want Subscriber)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				db: db.NewPgxMock(t),
			},
			args: args{
				ctx: context.Background(),
				arg: UpsertSubscriberParams{
					Email:  "test01@example.com",
					Active: true,
				},
			},
			want: Subscriber{
				ID:        uuid.New(),
				Email:     "test01@example.com",
				CreatedAt: pgtype.Timestamp{},
				UpdatedAt: pgtype.Timestamp{},
				Active:    true,
			},
			setupFn: func(mock pgxmock.PgxPoolIface, params UpsertSubscriberParams, want Subscriber) {
				rows := pgxmock.NewRows([]string{"id", "email", "created_at", "updated_at", "active"}).
					AddRow(want.ID, params.Email, want.CreatedAt, want.UpdatedAt, want.Active)
				mock.ExpectQuery(`INSERT (.+) CONFLICT\(email\)(.+) UPDATE`).WithArgs(params.Email, params.Active).WillReturnRows(rows)
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			mock := tt.fields.db.(pgxmock.PgxPoolIface)
			tt.setupFn(mock, tt.args.arg, tt.want)
			got, err := q.UpsertSubscriber(tt.args.ctx, tt.args.arg)
			if !tt.wantErr(t, err, fmt.Sprintf("UpsertSubscriber(%v, %v)", tt.args.ctx, tt.args.arg)) {
				return
			}
			assert.Equalf(t, tt.want, got, "UpsertSubscriber(%v, %v)", tt.args.ctx, tt.args.arg)
		})
	}
}
