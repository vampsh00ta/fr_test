package repository

import (
	"context"
	"fr/pkg/client"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	ClientRepository
	MessageRepository
	NewsletterRepository
	Tx(ctx context.Context) (pgx.Tx, error)
}
type Pg struct {
	client client.Client
}

func (pg Pg) Tx(ctx context.Context) (pgx.Tx, error) {
	tx, err := pg.client.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func New(client client.Client) Repository {
	return &Pg{
		client,
	}
}
