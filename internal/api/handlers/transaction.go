package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlesInit) CreateTransaction(ctx *gin.Context) {
	// id := ctx.Param("id")
	UserID := ctx.MustGet("userLogin").(int)
	// productID, _ := strconv.Atoi(id)

	err := h.service.CreateTransaction(ctx, UserID)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	successResponse(ctx, http.StatusOK, " success", nil)
}

func (h *handlesInit) GetTransactionHistory(ctx *gin.Context) {
	userID := ctx.MustGet("userLogin").(int)

	transactions, err := h.service.GetTransactionHistory(ctx, userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, http.StatusOK, "Transaction history retrieved successfully", transactions)
}
