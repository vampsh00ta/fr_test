package repository

import (
	"context"
	"fmt"
	"fr/internal/repository/models"
	"github.com/jackc/pgx/v5"
)

type ClientRepository interface {
	AddClient(ctx context.Context, tx pgx.Tx, client models.Client) (*models.Client, error)
	UpdateClientById(ctx context.Context, tx pgx.Tx, client models.Client) error
	DeleteClientById(ctx context.Context, tx pgx.Tx, id int) error
	GetClientsByFilter(ctx context.Context, tx pgx.Tx, filter models.Filter) ([]models.Client, error)
	GetClientsAll(ctx context.Context, tx pgx.Tx) ([]models.Client, error)
}

func (pg Pg) AddClient(ctx context.Context, tx pgx.Tx, client models.Client) (*models.Client, error) {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}
	q := `insert into client (tel_number, mobile_code , tag , timezone) 
		values ($1,$2,$3,$4) returning id
		 `
	if err = tx.QueryRow(ctx, q, client.TelNumber, client.MobileCode, client.Tag, client.Timezone).Scan(&client.Id); err != nil {
		return nil, err
	}
	return &client, nil
}

func (pg Pg) UpdateClientById(ctx context.Context, tx pgx.Tx, client models.Client) error {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}

	q := `
update client set tel_number  = $1 ,mobile_code = $2,tag = $3,timezone = $4 where id = $5
		 returning id`
	if err = tx.QueryRow(ctx, q, client.TelNumber, client.MobileCode, client.Tag, client.Timezone, client.Id).
		Scan(&client.Id); err != nil {

		return err
	}
	return nil
}

func (pg Pg) DeleteClientById(ctx context.Context, tx pgx.Tx, id int) error {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `
delete from  client where id = $1 returning id`
	if err = tx.QueryRow(ctx, q, id).Scan(&id); err != nil {

		return err
	}
	return nil
}

func (pg Pg) GetClientsByFilter(ctx context.Context, tx pgx.Tx, filter models.Filter) ([]models.Client, error) {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}

	q := `select * from  client where`

	input := make([]any, len(filter.Tag)+len(filter.MobileCodes))

	qTag := `tag in (`
	for i, tag := range filter.Tag {
		qTag += fmt.Sprintf("$%d,", i+1)
		input[i] = tag

	}
	qTag = qTag[0:len(qTag)-1] + ")"

	qCode := `mobile_code in (`

	for i, code := range filter.MobileCodes {
		qCode += fmt.Sprintf("$%d,", i+1+len(filter.Tag))
		input[i+len(filter.Tag)] = code

	}
	qCode = qCode[0:len(qCode)-1] + ")"

	q += qTag + ` or ` + qCode
	rows, err := tx.Query(ctx, q, input...)
	if err != nil {
		return nil, err
	}
	clients, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Client])
	if err != nil {
		return nil, err
	}
	return clients, nil
}
func (pg Pg) GetClientsAll(ctx context.Context, tx pgx.Tx) ([]models.Client, error) {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}

	q := `select * from  client`
	rows, err := tx.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	clients, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Client])
	if err != nil {
		fmt.Println("asdasd")
		return nil, err
	}
	return clients, nil
}
