package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlesInit) AddToCart(ctx *gin.Context) {
	id := ctx.Param("id")

	productID, _ := strconv.Atoi(id)
	userID := ctx.MustGet("userLogin").(int)

	err := h.service.AddToCart(ctx, productID, userID)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	successResponse(ctx, http.StatusOK, " success", productID)
}


func (h *handlesInit) GetCartByUserByID(ctx *gin.Context) {
	userID := ctx.MustGet("userLogin").(int)

	carts, err := h.service.GetCartByUserID(ctx, userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, http.StatusOK, "Success", carts)
}
