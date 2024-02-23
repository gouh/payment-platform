package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"payment-platform/internal/container"
	"payment-platform/internal/responses"
)

type (
	HealthHandlerInterface interface {
		HealthCheck(c *gin.Context)
	}
	HealthHandler struct {
		Db *pgxpool.Pool
	}
)

func (handler *HealthHandler) HealthCheck(c *gin.Context) {
	errPingDb := handler.Db.Ping(context.Background())
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
