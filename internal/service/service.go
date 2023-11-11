package service

import (
	"context"
	"fr/internal/repository"
	"github.com/go-co-op/gocron"
	"time"
)

type Service interface {
	Client
	Newsletter
	Scheduler
}
type service struct {
	repo      repository.Repository
	scheduler *gocron.Scheduler
}

func (s service) UpdateStartTimeNewsletter(ctx context.Context, id int, startTime time.Time) error {
	//TODO implement me
	panic("implement me")
}

func New(repo repository.Repository) Service {
	return &service{
		repo:      repo,
		scheduler: gocron.NewScheduler(time.UTC),
	}
}
