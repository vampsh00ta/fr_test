package v1

import "fr/internal/repository/models"

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
