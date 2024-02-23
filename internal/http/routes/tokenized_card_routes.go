package routes

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/handlers"
)

func SetupTokenizedCardRoutes(router *gin.RouterGroup, container *container.Container) {
	handler := handlers.NewTokenizedCardHandler(container)
	router.GET("/customers/:id/cards", handler.GetTokenizedCards)
	router.GET("/customers/:id/cards/:token", handler.GetTokenizedCardById)
	router.POST("/customers/:id/cards", handler.CreateTokenizedCard)
	router.DELETE("/customers/:id/cards/:token", handler.DeleteTokenizedCard)
}
