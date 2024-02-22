package http

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/middleware"
	"payment-platform/internal/http/routes"
)

func SetupRoutes(router *gin.Engine, container *container.Container) {

	// Grupo para la versión 1 de la API
	v1 := router.Group("/v1")
	{
		v1.Use(middleware.CorsMiddleware())

		// health routes
		routes.SetupHealthRoutes(v1, container)
	}
}
