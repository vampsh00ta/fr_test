package service

import (
	"context"
	"fr/internal/repository/models"
)

type Client interface {
	AddClient(ctx context.Context, client models.Client) (*models.Client, error)
	UpdateClientById(ctx context.Context, client models.Client) error
	DeleteClientById(ctx context.Context, id int) error
}

func (s service) AddClient(ctx context.Context, client models.Client) (*models.Client, error) {
	tx, err := s.repo.Tx(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.repo.AddClient(ctx, tx, client)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	return res, nil
}

func (s service) UpdateClientById(ctx context.Context, client models.Client) error {
	tx, err := s.repo.Tx(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		return err
	}
	err = s.repo.UpdateClientById(ctx, tx, client)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}

func (s service) DeleteClientById(ctx context.Context, id int) error {
	tx, err := s.repo.Tx(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		return err
	}
	err = s.repo.DeleteClientById(ctx, tx, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}
