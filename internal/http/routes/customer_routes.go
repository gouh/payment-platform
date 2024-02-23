package routes

import (
	"github.com/gin-gonic/gin"
	"payment-platform/internal/container"
	"payment-platform/internal/http/handlers"
)

func SetupCustomerRoutes(router *gin.RouterGroup, container *container.Container) {
	handler := handlers.NewCustomerHandler(container)
	router.GET("/customers", handler.GetCustomers)
	router.GET("/customers/:id", handler.GetCustomerById)
	router.POST("/customers", handler.CreateCustomer)
	router.PATCH("/customers/:id", handler.UpdateCustomer)
	router.DELETE("/customers/:id", handler.DeleteCustomer)
}
