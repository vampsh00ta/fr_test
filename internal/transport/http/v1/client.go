package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
