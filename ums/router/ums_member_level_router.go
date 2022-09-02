package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsMemberLevelRouter)
}

func registerUmsMemberLevelRouter(e *gin.Engine) {
	iService := service.UmsMemberLevelService{}
	iApi := api.UmsMemberLevelApi{
		Service: iService,
	}
	memberLevel := e.Group("/memberLevel")
	{
		memberLevel.GET("/list", iApi.List)
	}
}
