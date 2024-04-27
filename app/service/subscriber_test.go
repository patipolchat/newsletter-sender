package service

import (
	"context"
	"energy-response-assignment/app/repository"
	"energy-response-assignment/config"
	"energy-response-assignment/db/sqlc"
	"energy-response-assignment/entity"
	mockRepository "energy-response-assignment/mocks/app/repository"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNewSubscriber(t *testing.T) {
	type args struct {
		cfg            *config.Config
		subscriberRepo repository.Subscriber
	}
	tests := []struct {
		name string
		args args
		want Subscriber
	}{
		{
			name: "TestNewSubscriber",
			args: args{
				cfg:            &config.Config{},
				subscriberRepo: mockRepository.NewSubscriber(t),
			},
			want: &subscriberImpl{
				cfg:            &config.Config{},
				subscriberRepo: mockRepository.NewSubscriber(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewSubscriber(tt.args.cfg, tt.args.subscriberRepo), "NewSubscriber(%v, %v)", tt.args.cfg, tt.args.subscriberRepo)
		})
	}
}

func Test_subscriberImpl_Subscribe(t *testing.T) {
	type fields struct {
		cfg            *config.Config
		subscriberRepo repository.Subscriber
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *entity.Subscriber
		beforeFn func(args args, fields fields, want *entity.Subscriber)
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "Subscribe Test",
			fields: fields{
				cfg: &config.Config{
					Database: &config.Database{
						ReadTimeout: 1,
					},
				},
				subscriberRepo: mockRepository.NewSubscriber(t),
			},
			args: args{
				ctx:   context.Background(),
				email: "test@example.com",
			},
			want: &entity.Subscriber{
				ID:        uuid.New(),
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Active:    true,
			},
			beforeFn: func(args args, fields fields, want *entity.Subscriber) {
				createdAt := new(pgtype.Timestamp)
				err := createdAt.Scan(want.CreatedAt)
				assert.NoError(t, err)
				updatedAt := new(pgtype.Timestamp)
				err = updatedAt.Scan(want.UpdatedAt)
				assert.NoError(t, err)
				fields.subscriberRepo.(*mockRepository.Subscriber).EXPECT().
					UpsertSubscriber(mock.Anything, sqlc.UpsertSubscriberParams{Email: args.email, Active: true}).
					Return(sqlc.Subscriber{
						ID:        want.ID,
						Email:     want.Email,
						CreatedAt: *createdAt,
						UpdatedAt: *updatedAt,
						Active:    want.Active,
					}, nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "Subscribe Test",
			fields: fields{
				cfg: &config.Config{
					Database: &config.Database{
						ReadTimeout: 1,
					},
				},
				subscriberRepo: mockRepository.NewSubscriber(t),
			},
			args: args{
				ctx:   context.Background(),
				email: "test@example.com",
			},
			want: nil,
			beforeFn: func(args args, fields fields, want *entity.Subscriber) {
				fields.subscriberRepo.(*mockRepository.Subscriber).EXPECT().
					UpsertSubscriber(mock.Anything, sqlc.UpsertSubscriberParams{Email: args.email, Active: true}).
					Return(sqlc.Subscriber{}, fmt.Errorf("error"))
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &subscriberImpl{
				cfg:            tt.fields.cfg,
				subscriberRepo: tt.fields.subscriberRepo,
			}
			tt.beforeFn(tt.args, tt.fields, tt.want)
			got, err := s.Subscribe(tt.args.ctx, tt.args.email)
			tt.wantErr(t, err, fmt.Sprintf("Subscribe(%v, %v)", tt.args.ctx, tt.args.email))
			assert.Equalf(t, tt.want, got, "Subscribe(%v, %v)", tt.args.ctx, tt.args.email)
		})
	}
}

func Test_subscriberImpl_Unsubscribe(t *testing.T) {
	type fields struct {
		cfg            *config.Config
		subscriberRepo repository.Subscriber
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *entity.Subscriber
		beforeFn func(args args, fields fields, want *entity.Subscriber)
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "UnSubscribe Test",
			fields: fields{
				cfg: &config.Config{
					Database: &config.Database{
						ReadTimeout: 1,
					},
				},
				subscriberRepo: mockRepository.NewSubscriber(t),
			},
			args: args{
				ctx:   context.Background(),
				email: "test@example.com",
			},
			want: &entity.Subscriber{
				ID:        uuid.New(),
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Active:    true,
			},
			beforeFn: func(args args, fields fields, want *entity.Subscriber) {
				createdAt := new(pgtype.Timestamp)
				err := createdAt.Scan(want.CreatedAt)
				assert.NoError(t, err)
				updatedAt := new(pgtype.Timestamp)
				err = updatedAt.Scan(want.UpdatedAt)
				assert.NoError(t, err)
				fields.subscriberRepo.(*mockRepository.Subscriber).EXPECT().
					UpsertSubscriber(mock.Anything, sqlc.UpsertSubscriberParams{Email: args.email, Active: false}).
					Return(sqlc.Subscriber{
						ID:        want.ID,
						Email:     want.Email,
						CreatedAt: *createdAt,
						UpdatedAt: *updatedAt,
						Active:    want.Active,
					}, nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "UnSubscribe Test",
			fields: fields{
				cfg: &config.Config{
					Database: &config.Database{
						ReadTimeout: 1,
					},
				},
				subscriberRepo: mockRepository.NewSubscriber(t),
			},
			args: args{
				ctx:   context.Background(),
				email: "test@example.com",
			},
			want: nil,
			beforeFn: func(args args, fields fields, want *entity.Subscriber) {
				fields.subscriberRepo.(*mockRepository.Subscriber).EXPECT().
					UpsertSubscriber(mock.Anything, sqlc.UpsertSubscriberParams{Email: args.email, Active: false}).
					Return(sqlc.Subscriber{}, fmt.Errorf("error"))
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &subscriberImpl{
				cfg:            tt.fields.cfg,
				subscriberRepo: tt.fields.subscriberRepo,
			}
			tt.beforeFn(tt.args, tt.fields, tt.want)
			got, err := s.Unsubscribe(tt.args.ctx, tt.args.email)
			tt.wantErr(t, err, fmt.Sprintf("UnSubscribe(%v, %v)", tt.args.ctx, tt.args.email))
			assert.Equalf(t, tt.want, got, "UnSubscribe(%v, %v)", tt.args.ctx, tt.args.email)
		})
	}
}
