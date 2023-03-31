package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("TOKEN") != c.Request.Header.Get("Authentication") {
			c.AbortWithStatusJSON(
				400,
				gin.H{"error": "No auth"},
			)
			return
		}
	}
}
