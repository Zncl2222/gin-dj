package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(router *gin.Engine) {
	router.Use(tokenAuthHandler())
}

func tokenAuthHandler() gin.HandlerFunc {
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
