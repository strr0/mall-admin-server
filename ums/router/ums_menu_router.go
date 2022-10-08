package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsMenuRouter)
}

func registerUmsMenuRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.UmsMenuService{DB: db}
	iApi := api.UmsMenuApi{Service: iService}
	menu := e.Group("/menu")
	{
		menu.POST("/create", iApi.Create)
		menu.POST("/update/:id", iApi.Update)
		menu.GET("/:id", iApi.GetItem)
		menu.POST("/delete/:id", iApi.Delete)
		menu.GET("/list/:parentId", iApi.List)
		menu.GET("/treeList", iApi.TreeList)
		menu.POST("/updateHidden/:id", iApi.UpdateHidden)
	}
}
