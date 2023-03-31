package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LimitMiddleware(limit int) gin.HandlerFunc {
	// create a buffered channel with 2000 spaces
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
