package v1

//func (t transport) getMessageStatuses(ctx *gin.Context) {
//	idStr := ctx.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		t.logger.Error(err, "http-v1-getLastMessageStatuses")
//		errorResponse(ctx, http.StatusBadRequest, ValidationError)
//		return
//
//	}
//	msgs, err := t.service.GetLastStatusesByNewsletterId(ctx.Request.Context(), id, status)
//	if err != nil {
//		t.logger.Error(err, "http-v1-getLastMessageStatuses")
//		errorResponse(ctx, http.StatusInternalServerError, ServerError)
//		return
//
//	}
//	ctx.JSON(http.StatusOK, &getLastMessageStatusesResponse{msgs})
//
//}
