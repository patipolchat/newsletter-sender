package repository

import (
	"energy-response-assignment/config"
	"energy-response-assignment/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewSubscriber(t *testing.T) {
	type args struct {
		cfg  *config.Config
		pool *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want Subscriber
	}{
		{
			name: "success",
			args: args{
				cfg: &config.Config{},
			},
			want: &subscriberImpl{
				cfg:     &config.Config{},
				Queries: sqlc.New(&pgxpool.Pool{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSubscriber(tt.args.cfg, tt.args.pool)
			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(got))
		})
	}
}
