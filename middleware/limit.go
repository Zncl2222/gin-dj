package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LimitMiddleware(router *gin.Engine, limit int) {
	router.Use(limitHandler(limit))
}

func limitHandler(limit int) gin.HandlerFunc {
	semaphore := make(chan bool, limit)

	return func(c *gin.Context) {
		select {
		case semaphore <- true:
			c.Next()
			<-semaphore
		default:
			c.AbortWithStatusJSON(
				http.StatusServiceUnavailable,
				gin.H{"error": "Server is busy now, try later !"},
			)
			return
		}
	}
}
