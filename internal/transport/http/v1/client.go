package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary     Add
// @Description Создает запись клиента
// @Tags        Client
// @Accept      json
// @Param       username body addClientRequest true "username"
// @Produce     json
// @Success     201 {object} response
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /client/add [post]
func (t transport) addClient(ctx *gin.Context) {
	var request addClientRequest
	if err := ctx.BindJSON(&request); err != nil {
		t.logger.Error(err, "http-v1-addClient")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}
	res, err := t.service.AddClient(ctx.Request.Context(), request.Client)
	if err != nil {
		t.logger.Error(err, "http-v1-addClient")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusCreated, &res)

}

// @Summary     Update
// @Description Изменяет  запись клиента
// @Tags        Client
// @Accept      json
// @Param       username body updateClientRequest true "username"
// @Produce     json
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /client [put]
func (t transport) updateClient(ctx *gin.Context) {
	var request updateClientRequest
	if err := ctx.BindJSON(&request); err != nil {
		t.logger.Error(err, "http-v1-updateClient")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}
	err := t.service.UpdateClientById(ctx.Request.Context(), request.Client)
	if err != nil {
		t.logger.Error(err, "http-v1-updateClient")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusOK, &response{Status: "ok"})

}

// @Summary     Delete
// @Description Удаляет  запись клиента
// @Tags        Client
// @Accept      json
// @Param       username body deleteClientRequest true "username"
// @Produce     json
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /client [delete]
func (t transport) deleteClient(ctx *gin.Context) {
	var request deleteClientRequest
	if err := ctx.BindJSON(&request); err != nil {
		t.logger.Error(err, "http-v1-deleteClient")
		errorResponse(ctx, http.StatusBadRequest, ValidationError)
		return

	}
	err := t.service.DeleteClientById(ctx.Request.Context(), request.Id)
	if err != nil {
		t.logger.Error(err, "http-v1-deleteClient")
		errorResponse(ctx, http.StatusInternalServerError, ServerError)
		return

	}
	ctx.JSON(http.StatusOK, &response{Status: "ok"})

}

//func (t transport) deleteClient(ctx *gin.Context) {
//	var request addClientRequest
//	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
//		t.logger.Error(err, "http-v1-deleteClient")
//		errorResponse(ctx, http.StatusBadRequest, err)
//	}
//	if err := t.service.AddClient(ctx.Request.Context(), request.Client); err != nil {
//		t.logger.Error(err, "http-v1-deleteClient")
//		errorResponse(ctx, http.StatusInternalServerError, err)
//	}
//	ctx.JSON(201, response{Status: "ok"})
//}
//
//func (t transport) changeClient(ctx *gin.Context) {
//	var request addClientRequest
//	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
//		t.logger.Error(err, "http-v1-changeClient")
//		errorResponse(ctx, http.StatusBadRequest, err)
//	}
//	if err := t.service.AddClient(ctx.Request.Context(), request.Client); err != nil {
//		t.logger.Error(err, "http-v1-changeClient")
//		errorResponse(ctx, http.StatusInternalServerError, err)
//	}
//	ctx.JSON(201, response{Status: "ok"})
//}
