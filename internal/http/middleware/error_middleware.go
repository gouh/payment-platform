package middleware

import (
	"braces.dev/errtrace"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"payment-platform/internal/responses"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			logrus.WithFields(
				logrus.Fields{
					"error": errtrace.Errorf("%v", err),
					"trace": errtrace.FormatString(err),
				},
			).Error(errtrace.Errorf("%v", err))

			c.AbortWithStatusJSON(c.Writer.Status(), responses.GetErrorResponse(err.Error()))
		}
	}
}
