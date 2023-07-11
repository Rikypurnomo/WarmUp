package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/gin-gonic/gin"
)

func (h *handlesInit) FindCategory(ctx *gin.Context) {
	var categories models.Category

	if err := ctx.Bind(&categories); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	search := (ctx.Query("search"))

	category, pagination, err := h.service.ListCategory(ctx, page, limit, search)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	successResponsePagination(ctx, http.StatusOK, "succes", category, pagination)
}

func (h *handlesInit) CreateCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.Bind(&category); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.CreateCategory(ctx, &category)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	successResponse(ctx, http.StatusOK, " success", nil)
}
