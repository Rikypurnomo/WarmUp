package handlers

import (
	"net/http"

	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/gin-gonic/gin"
)

func (h *handlesInit) Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.Register(ctx, &user)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	successResponse(ctx, http.StatusOK, "register success", nil)
}

func (h *handlesInit) Login(ctx *gin.Context) {
	var loginRequest models.Login

	if err := ctx.Bind(&loginRequest); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	loginResponse, err := h.service.Login(ctx, &loginRequest)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	successResponse(ctx, http.StatusOK, "login success", loginResponse)
}
