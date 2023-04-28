package middleware

import (
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(router *gin.Engine, token string) {
	router.Use(tokenAuthHandler(token))
}

func tokenAuthHandler(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token != c.Request.Header.Get("Token-Authentication") {
			c.AbortWithStatusJSON(
				400,
				gin.H{"error": "No auth"},
			)
			return
		}
	}
}
