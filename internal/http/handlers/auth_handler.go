package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"payment-platform/config"
	"payment-platform/internal/container"
	"payment-platform/internal/requests"
	"payment-platform/internal/responses"
	"time"
)

type (
	AuthHandlerInterface interface {
		Login(c *gin.Context)
	}
	AuthHandler struct {
		AuthConfig config.AuthConfig
	}
)

func (handler *AuthHandler) Login(c *gin.Context) {
	var authDetails requests.AuthRequest
	if err := c.ShouldBindJSON(&authDetails); err != nil {
		c.JSON(http.StatusBadRequest, responses.GetErrorResponse("Invalid request"))
		return
	}
	if authDetails.Username != handler.AuthConfig.User || authDetails.Password != handler.AuthConfig.Password {
		c.JSON(http.StatusUnauthorized, responses.GetErrorResponse("Authentication failed"))
		return
	}

	claims := requests.CustomClaims{
		Username: authDetails.Username,
		Exp:      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(handler.AuthConfig.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func NewAuthHandler(container *container.Container) AuthHandlerInterface {
	return &AuthHandler{AuthConfig: *container.Config.AuthConfig}
}
