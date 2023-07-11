package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/gin-gonic/gin"
)

func (h *handlesInit) FindProduct(ctx *gin.Context) {
	var products models.Product

	if err := ctx.Bind(&products); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	search := (ctx.Query("search"))

	product, pagination, err := h.service.ListProduct(ctx, page, limit,search)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	successResponsePagination(ctx, http.StatusOK, "succes", product, pagination)
}

func (h *handlesInit) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	productID, _ := strconv.Atoi(id)

	err := h.service.DeleteProduct(ctx, productID)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, "Data not found!")
		return
	}
	successResponse(ctx, http.StatusOK, "succes", id)
}

func (h *handlesInit) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.Bind(&product); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.CreateProduct(ctx, &product)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	
	successResponse(ctx, http.StatusOK, " success", nil)
}
