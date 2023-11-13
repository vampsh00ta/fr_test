package v1

import (
	"fr/internal/repository/models"
	"github.com/gin-gonic/gin"
)

type response struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{Error: msg})
}

type createClientResponse struct {
	*models.Client
}

type createNewsletterResponse struct {
	Id int `json:"id"  binding:"required"  example:"1"`
}
type getLastMessageStatusesResponse struct {
	Messages []models.MessageStatus `json:"messages"  binding:"required"  `
}
type getNewsletter struct {
	*models.Newsletter
	Messages []models.MessageStatus `json:"messages"  binding:"required"  `
}
