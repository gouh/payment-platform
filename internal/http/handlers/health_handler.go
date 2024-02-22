package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-platform/internal/container"
	"payment-platform/internal/responses"
)

type (
	HealthHandlerInterface interface {
		HealthCheck(c *gin.Context)
	}
	HealthHandler struct {
		Db *sql.DB
	}
)

func (handler *HealthHandler) HealthCheck(c *gin.Context) {
	errPingDb := handler.Db.Ping()
	if errPingDb != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse("Error connecting to database"+errPingDb.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.GetHealthResponse())
}

func NewHealthHandler(container *container.Container) HealthHandlerInterface {
	return &HealthHandler{
		Db: container.Db,
	}
}
