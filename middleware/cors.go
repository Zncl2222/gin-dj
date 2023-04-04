package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(router *gin.Engine, config cors.Config) {
	router.Use(cors.New(config))
}
