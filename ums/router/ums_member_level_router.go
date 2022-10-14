package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsMemberLevelRouter)
}

func registerUmsMemberLevelRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.UmsMemberLevelService{DB: db}
	iApi := api.UmsMemberLevelApi{Service: iService}
	memberLevel := e.Group("/memberLevel").Use(config.AuthCheckRole("admin"))
	{
		memberLevel.GET("/list", iApi.List)
	}
}
