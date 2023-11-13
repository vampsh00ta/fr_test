package v1

import (
	"fr/internal/repository/models"
	"time"
)

type addClientRequest struct {
	models.Client
}
type createNewsletterRequest struct {
	models.Newsletter
}

type updateClientRequest struct {
	models.Client
}
type deleteClientRequest struct {
	Id int `json:"id"  binding:"required"  example:"1"`
}

type deleteNewsletterRequest struct {
	Id int `json:"id"  binding:"required"  example:"1"`
}
type updateNewsletterRequest struct {
	Id   int        `json:"id"  binding:"required"  example:"1"`
	Time *time.Time `json:"time" example:"1"`
	Text string     `json:"text" example:"1"`
}

type getLastMessageStatusesRequest struct {
	Id int `json:"id"  binding:"required"  example:"1"`
}
