package routes

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/handlers"
)

func SetupPaymentRoutes(router *gin.RouterGroup, container *container.Container) {
	handler := handlers.NewPaymentHandler(container)
	router.GET("/payments", handler.GetPayments)
	router.GET("/payments/:id", handler.GetPaymentById)
	router.POST("/payments", handler.CreatePayment)
}
