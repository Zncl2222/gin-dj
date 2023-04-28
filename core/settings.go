package core

import (
	"github.com/Zncl2222/gin-dj/middleware"
	"github.com/gin-contrib/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var DATABASE *gorm.DB
var DATABASE_URL string = "DB_URL"

func init() {
	var err error
	if DATABASE_URL != "DB_URL" {
		DATABASE, err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

		if err != nil {
			panic("Failed to connect database")
		}
	}
}

func GetDB() *gorm.DB {
	return DATABASE
}

// Middleware config
var CORSConfig = cors.Config{
	AllowOrigins:     []string{"https://foo.com"},
	AllowMethods:     []string{"PUT", "PATCH"},
	AllowHeaders:     []string{"Origin"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		return origin == "https://github.com"
	},
}

// Global midleware settings (for all APIs)
func GlobalMiddlewareInit(router *gin.Engine) {
	middleware.CORSMiddleware(router, CORSConfig)
	middleware.LimitMiddleware(router, 100)
	middleware.CSRFMiddleware(router)
}

// Custom midleware settings (for specific API)
func MiddlewareInit(router *gin.Engine) {
	// middleware.TokenAuthMiddleware(router, os.Getenv('yourtokenname'))
}
