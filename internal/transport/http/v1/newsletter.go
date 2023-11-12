package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary		CreateNewsletter
// @Description	Создает рассылку
// @Tags			Newsletter
// @Accept			json
// @Param  		username   body      createNewsletterRequest  true  "username"
// @Produce		json
// @Success		201 {object}    response
// @Failure		400	{object}	response
// @Failure		404	{object}	response
// @Failure		500	{object}	response
// @Router			/newsletter/create [post]
func (t transport) createNewsletter(ctx *gin.Context) {
	var request createNewsletterRequest
	if err := ctx.BindJSON(&request); err != nil {
		t.logger.Error(err, "http-v1-createNewsletter")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}
	newsletter, err := t.service.CreateNewsletter(ctx.Request.Context(), request.Newsletter)
	if err != nil {
		t.logger.Error(err, "http-v1-createNewsletter")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusCreated, &createNewsletterResponse{Id: newsletter.Id})

}

func (t transport) deleteNewsletter(ctx *gin.Context) {
	var request deleteNewsletterRequest
	if err := ctx.BindJSON(&request); err != nil {
		t.logger.Error(err, "http-v1-deleteNewsletter")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}

	err := t.service.DeleteNewsletter(ctx.Request.Context(), request.Id)
	if err != nil {
		t.logger.Error(err, "http-v1-deleteNewsletter")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusCreated, &response{Status: "ok"})

}
func (t transport) updateNewsletter(ctx *gin.Context) {
	var request updateNewsletterRequest
	if err := ctx.BindJSON(&request); err != nil {
		t.logger.Error(err, "http-v1-updateNewsletter")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}
	if request.Time == nil && request.Text == "" {
		t.logger.Error(ValidationError, "http-v1-updateNewsletter")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return
	}
	err := t.service.UpdateNewsletter(ctx.Request.Context(), request.Id, request.Time, request.Text)
	if err != nil {
		t.logger.Error(err, "http-v1-updateNewsletter")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusCreated, &response{Status: "ok"})

}
