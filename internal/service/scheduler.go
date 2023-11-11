package service

import (
	"context"
	"github.com/go-co-op/gocron"
	"strconv"
	"time"
)

type Scheduler interface {
	AddToScheduleNewsletter(ctx context.Context, newsletterId int, t *time.Time) error
}

//	type Scheduler interface {
//		AddToScheduleNewsletter(newsletterId int)
//	}
//
//	type scheduler struct {
//		gocron.Scheduler
//	}
func (s service) AddToScheduleNewsletter(ctx context.Context, newsletterId int, t *time.Time) error {

	job, err := s.scheduler.Every(1).Second().
		StartAt(*t).Tag(strconv.Itoa(newsletterId)).
		DoWithJobDetails(s.schedulerNewsletter, ctx, newsletterId)
	if err != nil {
		return job.Error()
	}
	s.scheduler.StartAsync()
	return nil
}
func (s service) schedulerNewsletter(ctx context.Context, newsletterId int, job gocron.Job) {

	err := s.SendNewsletter(ctx, newsletterId)
	if err == nil || job.RunCount() == 0 {
		s.scheduler.RemoveByTag(strconv.Itoa(newsletterId))
	}

}
