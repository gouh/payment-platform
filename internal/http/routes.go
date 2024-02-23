package http

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/middleware"
	"payment-platform/internal/http/routes"
)

func SetupRoutes(router *gin.Engine, container *container.Container) {
	v1 := router.Group("/v1")
	{
		v1.Use(middleware.ErrorMiddleware())
		v1.Use(middleware.RecoveryMiddleware())
		v1.Use(middleware.CorsMiddleware())

		// health routes
		routes.SetupHealthRoutes(v1, container)

		// merchant routes
		routes.SetupMerchantRoutes(v1, container)

		// customer routes
		routes.SetupCustomerRoutes(v1, container)

		// customer-card routes
		routes.SetupTokenizedCardRoutes(v1, container)

		// payment routes
		routes.SetupPaymentRoutes(v1, container)
	}
}
