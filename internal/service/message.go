package service

import (
	"context"
	"fr/internal/repository/models"
)

type Message interface {
	GetLastStatusesByNewsletterId(ctx context.Context, id int, status string) ([]models.MessageStatus, error)
	GetStatusesByMessageId(ctx context.Context, id int, status string) ([]models.MessageStatus, error)
}

func (s service) GetLastStatusesByNewsletterId(ctx context.Context, id int, status string) ([]models.MessageStatus, error) {
	tx, err := s.repo.Tx(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	msgs, err := s.repo.GetLastStatusesByNewsletterId(ctx, tx, id, status)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	return msgs, nil
}

func (s service) GetStatusesByMessageId(ctx context.Context, id int, status string) ([]models.MessageStatus, error) {
	tx, err := s.repo.Tx(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	msgs, err := s.repo.GetLastStatusesByNewsletterId(ctx, tx, id, status)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}
	return msgs, nil
}
