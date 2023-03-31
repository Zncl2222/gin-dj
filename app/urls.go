package app

import (
	"github.com/gin-gonic/gin"
)

func DefaultRegister(router *gin.RouterGroup) {
	router.GET("/", List)
	router.POST("/", Create)
	router.PATCH("/", Update)
	router.DELETE("/", Delete)
}
