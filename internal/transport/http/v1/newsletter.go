package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t transport) createNewsletter(ctx *gin.Context) {
	var request createNewsletterRequest
	if err := ctx.BindJSON(&request); err != nil {
		t.logger.Error(err, "http-v1-createNewsletter")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}

	err := t.service.CreateNewsletter(ctx.Request.Context(), request.Newsletter)
	if err != nil {
		t.logger.Error(err, "http-v1-createNewsletter")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusCreated, &response{Status: "ok"})

}
