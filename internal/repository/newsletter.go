package repository

import (
	"context"
	"errors"
	"fmt"
	"fr/internal/repository/models"
	"github.com/jackc/pgx/v5"
	"time"
)

type NewsletterRepository interface {
	AddNewsletter(ctx context.Context, tx pgx.Tx, newsletter models.Newsletter) (*models.Newsletter, error)
	UpdateNewsletterEndTime(ctx context.Context, tx pgx.Tx, id int, endTime *time.Time) error
	//GetNewsletterMsgClientById(ctx context.Context, tx pgx.Tx, id int) ([]models.MessageClient, error)
	UpdateNewsletterById(ctx context.Context, tx pgx.Tx, id int, t *time.Time, text string) error
	DeleteNewsletterById(ctx context.Context, tx pgx.Tx, id int) error
	GetNewsletterById(ctx context.Context, tx pgx.Tx, id int) (*models.Newsletter, error)
}

func (pg Pg) AddNewsletter(ctx context.Context, tx pgx.Tx, newsletter models.Newsletter) (*models.Newsletter, error) {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}
	q := `insert into newsletter (start_time , text) 
		values ($1,$2) returning id
		 `
	if err = tx.QueryRow(ctx, q, newsletter.StartTime, newsletter.Text).
		Scan(&newsletter.Id); err != nil {
		return nil, err
	}
	return &newsletter, nil
}
func (pg Pg) UpdateNewsletterEndTime(ctx context.Context, tx pgx.Tx, id int, endTime *time.Time) error {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `update  newsletter  set end_time = $1
		where id = $2 returning id
		 `
	if err = tx.QueryRow(ctx, q, endTime, id).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (pg Pg) UpdateNewsletterById(ctx context.Context, tx pgx.Tx, id int, t *time.Time, text string) error {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `update  newsletter  set
		 `
	if t == nil && text == "" {
		return errors.New("empty")
	}
	input := []any{}
	inputCount := 1
	if t != nil {
		input = append(input, t)
		q += fmt.Sprintf("start_time = $%d,", inputCount)
		inputCount += 1
	}
	if text != "" {
		input = append(input, text)
		q += fmt.Sprintf("text = $%d,", inputCount)
		inputCount += 1
	}
	input = append(input, id)
	q = fmt.Sprintf(q[0:len(q)-1]+" where id = $%d returning id", inputCount)
	if err = tx.QueryRow(ctx, q, input...).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (pg Pg) DeleteNewsletterById(ctx context.Context, tx pgx.Tx, id int) error {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `delete from  newsletter where id = $1 returning id `
	if err = tx.QueryRow(ctx, q, id).Scan(&id); err != nil {
		return err
	}
	return nil
}
func (pg Pg) GetNewsletterById(ctx context.Context, tx pgx.Tx, id int) (*models.Newsletter, error) {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}
	q := `select * from  newsletter where id = $1`
	var newsletter models.Newsletter
	if err = tx.QueryRow(ctx, q, id).Scan(&newsletter.Id, &newsletter.StartTime, &newsletter.EndTime, &newsletter.Text); err != nil {
		return nil, err
	}
	return &newsletter, nil
}
