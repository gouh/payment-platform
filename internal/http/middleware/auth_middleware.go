package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"payment-platform/internal/container"
	"payment-platform/internal/requests"
	"payment-platform/internal/responses"
	"strings"
)

func AuthMiddleware(container *container.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.GetErrorResponse("JWT is required"))
			return
		}

		newToken := strings.ReplaceAll(tokenString, "Bearer ", "")
		token, errParseClaims := jwt.ParseWithClaims(newToken, &requests.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(container.Config.AuthConfig.SecretKey), nil
		})

		if errParseClaims != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.GetErrorResponse(errParseClaims.Error()))
			return
		}

		if claims, ok := token.Claims.(*requests.CustomClaims); ok && token.Valid {
			c.Set("username", claims.Username)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.GetErrorResponse("Invalid token"))
			return
		}

		c.Next()
	}
}
