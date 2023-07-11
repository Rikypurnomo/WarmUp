package server

import (
	"github.com/Rikypurnomo/warmup/internal/api/handlers"
	"github.com/Rikypurnomo/warmup/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func RouterApi(router *gin.Engine) *gin.Engine {
	// initiate handler
	h := handlers.InitiateHandlersInterface()

	router.GET("/", h.Welcome)
	router.GET("/balance", middleware.CheckServicesAdapter(), h.MyBalance)
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
	router.GET("/products", middleware.Auth(), h.FindProduct)
	router.POST("/product", middleware.Auth(), h.CreateProduct)
	router.DELETE("/product/:id", middleware.Auth(), h.DeleteProduct)
	router.POST("/category", middleware.Auth(), h.CreateCategory)
	router.GET("categories", h.FindCategory)
	router.POST("/cart/:id", middleware.Auth(), h.AddToCart)
	router.POST("/transaction", middleware.Auth(), h.CreateTransaction)
	router.GET("carts", middleware.Auth(), h.GetCartByUserByID)
	router.GET("/history", middleware.Auth(), h.GetTransactionHistory)

	return router
}
