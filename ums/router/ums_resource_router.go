package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsResourceRouter)
}

func registerUmsResourceRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.UmsResourceService{DB: db}
	iApi := api.UmsResourceApi{Service: iService}
	resource := e.Group("/resource")
	{
		resource.POST("/create", iApi.Create)
		resource.POST("/update/:id", iApi.Update)
		resource.GET("/:id", iApi.GetItem)
		resource.POST("/delete/:id", iApi.Delete)
		resource.GET("/list", iApi.List)
		resource.GET("/listAll", iApi.ListAll)
	}
}
