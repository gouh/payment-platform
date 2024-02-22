package routes

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/handlers"
)

func SetupHealthRoutes(router *gin.RouterGroup, container *container.Container) {
	handler := handlers.NewHealthHandler(container)
	router.GET("/health", handler.HealthCheck)
}
