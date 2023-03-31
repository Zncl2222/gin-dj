package gin_dj

import (
	"github.com/Zncl2222/gin-dj/app"
	"github.com/gin-gonic/gin"
)

func RouterInit(router *gin.Engine) {
	todos_group := router.Group("api/v1/todos")
	{
		app.DefaultRegister(todos_group)
	}

}
