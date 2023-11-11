package service

import "context"

type Email interface {
}

func (s service) SendEmail(ctx context.Context, id int) error {
	return nil
}
