package models

import (
	"time"
)

type Client struct {
	Id         int    `json:"id" db:"id" `
	TelNumber  string `json:"tel_number,omitempty" db:"tel_number,null" binding:"required"  example:"9999999999"`
	MobileCode string `json:"mobile_code,omitempty" db:"mobile_code,null" binding:"required"  example:"7"`
	Tag        string `json:"tag" db:"tag,null" binding:"required"  example:"tag"`
	Timezone   string `json:"timezone" db:"timezone,null" binding:"required"  example:"UTC"`
}

const (
	Created      = "Сообщение создано"
	Send         = "Сообщение отправлено"
	SendApiError = "Ошибка"
)

type Message struct {
	Id           int        `json:"id" db:"id"`
	SendTime     *time.Time `json:"send_time" db:"send_time" `
	CreationTime *time.Time `json:"creation_time" db:"creation_time"`
	Status       string     `json:"status" db:"status"  example:"status"`
	NewsletterId int        `json:"newsletter_id" db:"newsletter_id" `
	ClientId     int        `json:"client_id" db:"client_id"`
}

type Newsletter struct {
	Id        int        `json:"id" db:"id"`
	StartTime *time.Time `json:"start_time" db:"start_time" binding:"required"`
	EndTime   *time.Time `db:"end_time"`
	Text      string     `json:"text" db:"text" binding:"required"`
	Messages  []Message
	Filter    Filter `json:"filter"`
}
type Filter struct {
	MobileCodes []string `json:"mobile_code,omitempty" db:"mobile_code" "  example:"7"`
	Tag         []string `json:"tag" db:"id"  example:"tag"`
}

type MessageClient struct {
	MessageId    int        `json:"message_id" db:"message_id"`
	SendTime     *time.Time `json:"send_time" db:"send_time" `
	CreationTime *time.Time `json:"creation_time" db:"creation_time"`
	Status       string     `json:"status" db:"status"  example:"status"`
	NewsletterId int        `json:"newsletter_id" db:"newsletter_id" `
	ClientId     int        `json:"client_id" db:"client_id"`
	TelNumber    string     `json:"tel_number,omitempty" db:"tel_number" `
	MobileCode   string     `json:"mobile_code,omitempty" db:"mobile_code"`
	Tag          string     `json:"tag" db:"tag" `
	Timezone     string     `json:"timezone" db:"timezone"`
	Text         string     `json:"text" db:"text" `
}
