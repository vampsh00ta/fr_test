package service

import (
	"context"
	"errors"
	"fr/internal/repository/models"
	"time"
)

type Newsletter interface {
	CreateNewsletter(ctx context.Context, newsletter models.Newsletter) error
	SendNewsletter(ctx context.Context, id int) error
	UpdateStartTimeNewsletter(ctx context.Context, id int, startTime time.Time) error
	UpdateNewsletter(ctx context.Context, id int, newsletter models.Newsletter) error
}

func (s service) CreateNewsletter(ctx context.Context, newsletter models.Newsletter) error {
	var err error

	tx, err := s.repo.Tx(ctx)
	var clients []models.Client
	if newsletter.Filter.Tag == nil && newsletter.Filter.MobileCodes == nil {
		clients, err = s.repo.GetClientsAll(ctx, tx)
	} else {
		clients, err = s.repo.GetClientsByFilter(ctx, tx, newsletter.Filter)

	}
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	newsl, err := s.repo.AddNewsletter(ctx, tx, newsletter)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	messages := make([]models.Message, len(clients))
	for i, client := range clients {
		messages[i] = models.Message{NewsletterId: newsl.Id, ClientId: client.Id}
	}
	_, err = s.repo.AddMessages(ctx, tx, messages...)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	tx.Commit(ctx)
	if err := s.AddToScheduleNewsletter(context.Background(), newsl.Id, newsletter.StartTime); err != nil {
		return err
	}

	return nil
}

//добавить случай если апи вернуло ошибку

func (s service) SendNewsletter(ctx context.Context, id int) error {
	tx, err := s.repo.Tx(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		return err
	}
	msgs, err := s.repo.GetFullMsgsByNewsletterID(ctx, tx, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	msg_to_send := make(chan models.MessageClient, len(msgs))
	result := make(chan models.MessageClient, len(msgs))
	allDone := true
	for w := 1; w <= 10; w++ {
		go func(id int, msg_to_send <-chan models.MessageClient, result chan<- models.MessageClient) {
			for j := range msg_to_send {
				resp, err := sendMsg(j)
				if err != nil || resp.Code != 0 {
					allDone = false
					j.Status = models.SendApiError

					result <- j
					continue
				}

				t := time.Now()
				j.SendTime = &t
				j.Status = models.Send
				result <- j

			}
		}(w, msg_to_send, result)
	}

	for _, msg := range msgs {
		msg_to_send <- msg
	}
	close(msg_to_send)

	for a := 1; a <= len(msgs); a++ {
		msg := <-result
		if err := s.repo.UpdateMessageById(ctx, tx, msg.MessageId, msg.Status, msg.SendTime); err != nil {

			return err
		}
	}

	if allDone == false {
		tx.Commit(ctx)
		return errors.New("not_all_done")
	}
	endTime := time.Now()
	if err := s.repo.UpdateNewsletterEndTime(ctx, tx, id, &endTime); err != nil {
		return err
	}
	return nil
}
func (s service) UpdateNewsletter(ctx context.Context, id int, newsletter models.Newsletter) error {
	//var err error
	//
	//tx, err := s.repo.Tx(ctx)
	//var clients []models.Client
	//
	//if err != nil {
	//	tx.Rollback(ctx)
	//	return err
	//}
	//err = s.repo.UpdateNewsletterById(ctx, tx, id, newsletter.EndTime, newsletter.Text)
	//if err != nil {
	//	tx.Rollback(ctx)
	//	return err
	//}
	//messages := make([]models.Message, len(clients))
	//for i, client := range clients {
	//	messages[i] = models.Message{NewsletterId: newsl.Id, ClientId: client.Id}
	//}
	//_, err = s.repo.AddMessages(ctx, tx, messages...)
	//if err != nil {
	//	tx.Rollback(ctx)
	//	return err
	//}
	//tx.Commit(ctx)
	//if err := s.AddToScheduleNewsletter(context.Background(), newsl.Id, newsletter.StartTime); err != nil {
	//	return err
	//}

	return nil
}
