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
	GetFullMsgsByNewsletterID(ctx context.Context, tx pgx.Tx, id int) ([]models.MessageFull, error)
	UpdateMessageById(ctx context.Context, tx pgx.Tx, id int, status string, sendTime *time.Time) error
	AddMessageStatusById(ctx context.Context, tx pgx.Tx, id int, status string, t *time.Time) error
	AddMessagesStatus(ctx context.Context, tx pgx.Tx, messages ...models.Message) error
	DeleteMessagesByNewsletterId(ctx context.Context, tx pgx.Tx, id int) error
	UpdateStatusByNewsletterId(ctx context.Context, tx pgx.Tx, id int, t *time.Time, text string) error
	UpdateMessageStatuses(ctx context.Context, tx pgx.Tx, status string, messages ...models.Message) error
	GetMsgsByNewsletterID(ctx context.Context, tx pgx.Tx, id int) ([]models.Message, error)
	GetLastStatusesByNewsletterId(ctx context.Context, tx pgx.Tx, id int, status string) ([]models.MessageStatus, error)
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

func (pg Pg) GetFullMsgsByNewsletterID(ctx context.Context, tx pgx.Tx, id int) ([]models.MessageFull, error) {
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
	fullMsgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MessageFull])
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

	if err = pg.AddMessageStatusById(ctx, tx, id, status, t); err != nil {
		return err
	}

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

func (pg Pg) DeleteMessagesByNewsletterId(ctx context.Context, tx pgx.Tx, id int) error {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `delete from  message where newsletter_id = $1 returning id `
	if err = tx.QueryRow(ctx, q, id).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (pg Pg) UpdateStatusByNewsletterId(ctx context.Context, tx pgx.Tx, id int, t *time.Time, text string) error {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	q := `update  status set time = $1,text  = $2 where id = $3 returning`
	if err = tx.QueryRow(ctx, q, t, text, id).Scan(&id); err != nil {
		return err
	}
	return nil
}
func (pg Pg) GetMsgsByNewsletterID(ctx context.Context, tx pgx.Tx, id int) ([]models.Message, error) {
	var err error
	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}
	q := `select message.id ,newsletter_id,send_time,creation_time,client_id

          from (select * from newsletter where id = $1) as  newsletter
		  join  message  on message.newsletter_id = newsletter.id
		  join (select * from status where text != 'Сообщение отправлено')  as status on   message.id = status.message_id
		`
	rows, err := tx.Query(ctx, q, id)

	if err != nil {
		return nil, err
	}
	msgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Message])
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
func (pg Pg) UpdateMessageStatuses(ctx context.Context, tx pgx.Tx, status string, messages ...models.Message) error {
	var err error

	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return err
		}
	}
	t := time.Now()
	q := `insert into  status (text,time,message_id) values
		 `

	input := []any{status, &t}
	for i, message := range messages {
		str := fmt.Sprintf("($1,$2,$%d),", i+3)
		q += str

		input = append(input, message.Id)

	}

	q = q[0:len(q)-1] + " returning text"
	if err = tx.QueryRow(ctx, q, input...).Scan(&status); err != nil {
		return err
	}

	return nil
}

func (pg Pg) GetLastStatusesByNewsletterId(ctx context.Context, tx pgx.Tx, id int, status string) ([]models.MessageStatus, error) {
	var err error

	if tx == nil {
		tx, err = pg.client.Begin(ctx)
		if err != nil {
			return nil, err
		}
	}
	input := []any{id}
	q := `  select message.*,text,time from message
			join (
				select * from (
					SELECT *, ROW_NUMBER() OVER (PARTITION BY message_id ORDER BY time DESC
					) AS rank_no FROM status    ORDER BY message_id desc) 
				as rd where rd.rank_no = 1) 
				as status on status.message_id = message.id 
			

		 `
	if status != "" {
		q += `where newsletter_id = $1 and text = $2 order by newsletter_id desc`
		input = append(input, status)
	} else {
		q += `where newsletter_id = $1  order by newsletter_id desc`
	}
	rows, err := tx.Query(ctx, q, input...)

	if err != nil {
		return nil, err
	}
	msgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MessageStatus])
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
