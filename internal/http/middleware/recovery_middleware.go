package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"payment-platform/internal/responses"
	"runtime/debug"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errObj error
				if e, ok := err.(error); ok {
					errObj = e
				} else {
					errObj = errors.Errorf("%v", err)
				}

				stackTrace := string(debug.Stack())

				logData := map[string]interface{}{
					"error": errObj.Error(),
					"msg":   "Panic recovered",
					"trace": stackTrace,
				}

				//prettyJSON, jsonErr := json.MarshalIndent(logData, "", "   ")
				//if jsonErr != nil {
				//	logrus.WithField("error", jsonErr).Error("Error al formatear log de error")
				//} else {
				//}
				logrus.Error(logData)

				c.AbortWithStatusJSON(http.StatusInternalServerError, responses.GetErrorResponse(errObj.Error()))
			}
		}()
		c.Next()
	}
}
