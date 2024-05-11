package routes

import (
	"net/http"

	"github.com/Erwin011895/TaskManagementApp/handler"
	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/gin-gonic/gin"
)

func AuthRequired(h handler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, pkg.Response{
				Message: http.StatusText(http.StatusUnauthorized),
			})
		}
	}
}
