package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlesInit) Welcome(ctx *gin.Context) {
	successResponse(ctx, http.StatusOK, "Welcome to kiyora", nil)
}

func (h *handlesInit) MyBalance(ctx *gin.Context) {
	defer starSpan(ctx.Request.Context(), "MyBalance").End()

	res, err := h.service.GetBalance(ctx.Request.Context())
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	successResponse(ctx, http.StatusOK, "Get Data Balance", res)
}
