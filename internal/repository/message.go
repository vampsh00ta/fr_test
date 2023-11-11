package repository

import (
	"context"
	"fmt"
	"fr/internal/repository/models"
	"github.com/jackc/pgx/v5"
	"time"
)

type MessageRepository interface {
	AddMessage(ctx context.Context, tx pgx.Tx, message models.Message) (*models.Message, error)
	AddMessages(ctx context.Context, tx pgx.Tx, messages ...models.Message) (*[]models.Message, error)
	GetFullMsgsByNewsletterID(ctx context.Context, tx pgx.Tx, id int) ([]models.MessageClient, error)
	UpdateMessageById(ctx context.Context, tx pgx.Tx, id int, status string, sendTime *time.Time) error
	AddMessageStatusById(ctx context.Context, tx pgx.Tx, id int, status string, t *time.Time) error
	AddMessagesStatus(ctx context.Context, tx pgx.Tx, messages ...models.Message) error

	//SetMessageDone(ctx context.Context, tx pgx.Tx, id int, status string, sendTime *time.Time) error
}

func (pg Pg) AddMessage(ctx context.Context, tx pgx.Tx, message models.Message) (*models.Message, error) {
	var err error

	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}
	q := `insert into message (creation_time, status , newsletter_id , client_id) 
		values ($1,$2,$3,$4) returning id
		 `
	if err = tx.QueryRow(ctx, q, time.Now(),
		models.Created,
		message.NewsletterId,
		message.ClientId).
		Scan(&message.Id); err != nil {
		return nil, err
	}
	return &message, nil
}
func (pg Pg) AddMessages(ctx context.Context, tx pgx.Tx, messages ...models.Message) (*[]models.Message, error) {
	var err error

	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}

	q := `insert into message (creation_time , newsletter_id , client_id) 
		values 
		 `
	input := []any{time.Now()}
	for i, message := range messages {
		str := fmt.Sprintf("($1,$%d,$%d),", i*2+2, i*2+3)
		q += str

		input = append(input, message.NewsletterId, message.ClientId)

	}

	q = q[0:len(q)-1] + " returning *"
	rows, err := tx.Query(ctx, q, input...)
	if err != nil {
		return nil, err
	}
	msgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Message])
	if err != nil {
		return nil, err
	}
	for i := range msgs {
		msgs[i].Status.Text = models.Created

	}
	if err := pg.AddMessagesStatus(ctx, tx, msgs...); err != nil {
		return nil, err
	}

	return &msgs, nil
}

func (pg Pg) GetFullMsgsByNewsletterID(ctx context.Context, tx pgx.Tx, id int) ([]models.MessageClient, error) {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}
	q := `select message.id as message_id, newsletter_id,send_time,creation_time,client_id,tel_number,mobile_code,tag,timezone,newsletter.text

          from (select * from newsletter where id = $1) as  newsletter
		  join  message  on message.newsletter_id = newsletter.id
		  join (select * from status where text != 'Сообщение отправлено')  as status on   message.id = status.message_id
		  join client on client.id  = message.client_id`
	rows, err := tx.Query(ctx, q, id)

	if err != nil {
		return nil, err
	}
	fullMsgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MessageClient])
	if err != nil {
		return nil, err
	}

	return fullMsgs, nil
}
func (pg Pg) UpdateMessageById(ctx context.Context, tx pgx.Tx, id int, status string, sendTime *time.Time) error {
	var err error

	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `update  message  set send_time = $1 where id = $2 returning id`

	if err = tx.QueryRow(ctx, q, sendTime, id).Scan(&id); err != nil {
		return err
	}

	var t *time.Time
	if t == nil {
		_t := time.Now()
		t = &_t

	}

	//q1 := `insert  into  status (time,text,message_id) values ($1,$2,$3)`
	//
	//if err = tx.QueryRow(ctx, q1, t, status, id).Scan(&id); err != nil {
	//	return err
	//}
	//if err = pg.AddMessageStatusById(ctx, tx, id, status, t); err != nil {
	//	return err
	//}

	return err
}
func (pg Pg) AddMessageStatusById(ctx context.Context, tx pgx.Tx, id int, status string, t *time.Time) error {
	var err error

	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}

	q := `insert  into  status (time,text,message_id) values ($1,$2,$3) returning id`

	if err = tx.QueryRow(ctx, q, t, status, id).Scan(&id); err != nil {
		return err
	}
	return err

}
func (pg Pg) AddMessagesStatus(ctx context.Context, tx pgx.Tx, messages ...models.Message) error {
	var err error

	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `insert  into  status (time,text,message_id) values
		 `
	input := []any{time.Now()}
	for i, message := range messages {
		str := fmt.Sprintf("($1,$%d,$%d),", i*2+2, i*2+3)
		q += str
		input = append(input, message.Status.Text, message.Id)

	}
	fmt.Println(q)
	q = q[0:len(q)-1] + " returning id"
	var id int
	if err = tx.QueryRow(ctx, q, input...).Scan(&id); err != nil {
		return err
	}

	return err

}
