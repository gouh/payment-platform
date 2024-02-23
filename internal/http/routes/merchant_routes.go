package routes

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/handlers"
)

func SetupMerchantRoutes(router *gin.RouterGroup, container *container.Container) {
	handler := handlers.NewMerchantHandler(container)
	router.GET("/merchants", handler.GetMerchants)
	router.GET("/merchants/:id", handler.GetMerchantById)
	router.POST("/merchants", handler.CreateMerchant)
	router.PATCH("/merchants/:id", handler.UpdateMerchant)
	router.DELETE("/merchants/:id", handler.DeleteMerchant)
}
