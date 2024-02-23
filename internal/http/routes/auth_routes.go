package routes

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/handlers"
)

func SetupAuthRoutes(router *gin.RouterGroup, container *container.Container) {
	handler := handlers.NewAuthHandler(container)
	router.POST("/login", handler.Login)
}
