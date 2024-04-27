package service

import (
	"context"
	"energy-response-assignment/app/repository"
	"energy-response-assignment/config"
	"energy-response-assignment/db/sqlc"
	"energy-response-assignment/entity"
	"time"
)

type Subscriber interface {
	Subscribe(ctx context.Context, email string) (*entity.Subscriber, error)
	Unsubscribe(ctx context.Context, email string) (*entity.Subscriber, error)
}

type subscriberImpl struct {
	cfg            *config.Config
	subscriberRepo repository.Subscriber
}

func (s *subscriberImpl) Subscribe(ctx context.Context, email string) (*entity.Subscriber, error) {
	return s.upsertSubscriber(ctx, email, true)
}

func (s *subscriberImpl) Unsubscribe(ctx context.Context, email string) (*entity.Subscriber, error) {
	return s.upsertSubscriber(ctx, email, false)
}

func (s *subscriberImpl) upsertSubscriber(ctx context.Context, email string, active bool) (*entity.Subscriber, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.Database.ReadTimeout*time.Second)
	defer cancel()
	result, err := s.subscriberRepo.UpsertSubscriber(ctx, sqlc.UpsertSubscriberParams{Email: email, Active: active})
	if err != nil {
		return nil, err
	}
	return &entity.Subscriber{
		ID:        result.ID,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
		Active:    result.Active,
	}, nil
}

func NewSubscriber(cfg *config.Config, subscriberRepo repository.Subscriber) Subscriber {
	return &subscriberImpl{
		cfg:            cfg,
		subscriberRepo: subscriberRepo,
	}
}
