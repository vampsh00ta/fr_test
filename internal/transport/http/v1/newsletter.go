package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary     Create
// @Description Создает рассылку.Если убрать фильтр, то добавит в рассылку всех клиентов.Можно добавить только 1 из фильтров.
// @Tags        Newsletter
// @Accept      json
// @Param       username body createNewsletterRequest true "body"
// @Produce     json
// @Success     201 {object} response
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /newsletter/create [post]
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

// @Summary     Delete
// @Description Удаляет рассылку.
// @Tags        Newsletter
// @Accept      json
// @Param       username body deleteNewsletterRequest true "body"
// @Produce     json
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /newsletter [delete]
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
	ctx.JSON(http.StatusOK, &response{Status: "ok"})

}

// @Summary     Update
// @Description Изменяет рассылку
// @Tags        Newsletter
// @Accept      json
// @Param       username body updateNewsletterRequest true "body"
// @Produce     json
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /newsletter [patch]
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

// @Summary     GetLastMessageStatuses
// @Description Выводит последние статусы сообщения по id рассылки
// @Tags        Newsletter
// @Accept      json
// @Param       id   path      int  true  "Newsletter ID"
// @Param status query string false "status"
// @Produce     json
// @Success     200 {object} getLastMessageStatusesResponse
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /newsletter/{id}/messages [get]
func (t transport) getLastMessageStatuses(ctx *gin.Context) {
	idStr := ctx.Param("id")
	status := ctx.Query("status")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		t.logger.Error(err, "http-v1-getLastMessageStatuses")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}
	msgs, err := t.service.GetLastStatusesByNewsletterId(ctx.Request.Context(), id, status)
	if err != nil {
		t.logger.Error(err, "http-v1-getLastMessageStatuses")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusOK, &getLastMessageStatusesResponse{msgs})

}

// @Summary     GetNewsletter
// @Description Выводит информацию о рассылке
// @Tags        Newsletter
// @Accept      json
// @Param       id   path      int  true  "Newsletter ID"
// @Produce     json
// @Success     200 {object} getNewsletter
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /newsletter/{id} [get]
func (t transport) getNewsletter(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		t.logger.Error(err, "http-v1-getLastMessageStatuses")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}
	newsletter, err := t.service.GetNewsletter(ctx.Request.Context(), id)
	if err != nil {
		t.logger.Error(err, "http-v1-getLastMessageStatuses")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusOK, &getNewsletter{newsletter, newsletter.Messages})

}
