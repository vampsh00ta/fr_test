package service

import (
	"context"
	"fmt"
	"fr/internal/repository/models"
)

type Message interface {
	GetLastStatusesByNewsletterId(ctx context.Context, id int, status string) ([]models.MessageStatus, error)
}

func (s service) GetLastStatusesByNewsletterId(ctx context.Context, id int, status string) ([]models.MessageStatus, error) {
	tx, err := s.repo.Tx(ctx)
	fmt.Println(tx, ctx)
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
