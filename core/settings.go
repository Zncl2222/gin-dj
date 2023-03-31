package core

import (
	"github.com/Zncl2222/gin-dj/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalMiddlewareHandlers = []interface{}{
	middleware.TokenAuthHandler,
	middleware.CORSMiddleware,
}

var DATABASE *gorm.DB
var DATABASE_URL string = "DB_URL"

func init() {
	var err error
	DATABASE, err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return DATABASE
}

func MiddlewareInit(router *gin.Engine, router_group_name ...string) {
	for _, handler := range GlobalMiddlewareHandlers {
		// regist global middleware for all APIs
		router.Use(handler.(func() gin.HandlerFunc)())
	}

	if router_group_name != nil {
		// regist custom middleware for specific APIs only
		router_group := router.Group(router_group_name[0])
		router_group.Use(middleware.LimitMiddleware(5))
	}
}
