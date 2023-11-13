package models

import (
	"time"
)

const (
	Created           = "Сообщение создано"
	Send              = "Сообщение отправлено"
	NewsletterDeleted = "Рассылка была удалена"
	TimeChanged       = "Время рассылки изменено"
	ParametersChanged = "Параметры рассылки изменены"
	SendApiError      = "Ошибка"
)

type Client struct {
	Id         int    `json:"id" db:"id" `
	TelNumber  int    `json:"tel_number,omitempty" db:"tel_number,null" binding:"required"  example:"9999999999"`
	MobileCode string `json:"mobile_code,omitempty" db:"mobile_code,null" binding:"required"  example:"7"`
	Tag        string `json:"tag" db:"tag,null" binding:"required"  example:"tag"`
	Timezone   string `json:"timezone" db:"timezone,null" binding:"required"  example:"UTC"`
}

type Message struct {
	Id           int        `json:"id,omitempty" db:"id,omitempty"`
	SendTime     *time.Time `json:"send_time,omitempty" db:"send_time,omitempty" `
	CreationTime *time.Time `json:"creation_time,omitempty" db:"creation_time,omitempty"`
	Status       Status     `json:"status,omitempty" db:"-"  `
	NewsletterId int        `json:"newsletter_id,omitempty" db:"newsletter_id,omitempty" `
	ClientId     int        `json:"client_id,omitempty" db:"client_id,omitempty"`
}
type Status struct {
	Id   int        `json:"id,omitempty" db:"id,omitempty"`
	Text string     `json:"text,omitempty" db:"text,omitempty"  example:"status,omitempty"`
	Time *time.Time `json:"time,omitempty" db:"time,omitempty" `
}
type Newsletter struct {
	Id        int             `json:"-" db:"id"`
	StartTime *time.Time      `json:"start_time" db:"start_time" binding:"required" example:"2023-11-12T16:45:00Z"`
	EndTime   *time.Time      `json:"-" db:"end_time"`
	Text      string          `json:"text" db:"text" binding:"required"`
	Messages  []MessageStatus `json:"-"`
	Filter    Filter          `json:"filter" db:"filter"`
}
type Filter struct {
	MobileCodes []string `json:"mobile_code,omitempty" db:"mobile_code,omitempty" example:"test"`
	Tag         []string `json:"tag,omitempty" db:"id,omitempty"  example:"tag"`
}

type MessageFull struct {
	MessageId    int        `json:"message_id" db:"message_id"`
	SendTime     *time.Time `json:"send_time" db:"send_time" `
	CreationTime *time.Time `json:"creation_time" db:"creation_time"`
	Status       string     `json:"status" db:"-"  example:"status"`
	NewsletterId int        `json:"newsletter_id" db:"newsletter_id" `
	ClientId     int        `json:"client_id" db:"client_id"`
	TelNumber    int        `json:"tel_number,omitempty" db:"tel_number" `
	MobileCode   string     `json:"mobile_code,omitempty" db:"mobile_code"`
	Tag          string     `json:"tag" db:"tag" `
	Timezone     string     `json:"timezone" db:"timezone"`
	Text         string     `json:"text" db:"text" `
}

type MessageStatus struct {
	Id           int        `json:"id" db:"id"`
	SendTime     *time.Time `json:"send_time" db:"send_time" `
	CreationTime *time.Time `json:"creation_time" db:"creation_time"`
	NewsletterId int        `json:"newsletter_id" db:"newsletter_id" `
	ClientId     int        `json:"client_id" db:"client_id"`
	Status       string     `json:"status" db:"text"  example:"status"`
	Time         *time.Time `json:"time" db:"time" `
}
