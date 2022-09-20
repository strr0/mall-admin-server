package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/api"
	"mall-admin-server/pms/service"
)

func init() {
	routers = append(routers, registerPmsProductAttributeCategoryRouter)
}

func registerPmsProductAttributeCategoryRouter(e *gin.Engine) {
	iService := service.PmsProductAttributeCategoryService{}
	iApi := api.PmsProductAttributeCategoryApi{
		Service: iService,
	}
	attributeCategory := e.Group("/productAttribute/category")
	{
		attributeCategory.POST("/create", iApi.Create)
		attributeCategory.POST("/update/:id", iApi.Update)
		attributeCategory.POST("/delete/:id", iApi.Delete)
		attributeCategory.GET("/:id", iApi.GetItem)
		attributeCategory.GET("/list", iApi.List)
		attributeCategory.GET("/list/withAttr", iApi.ListWithAttr)
	}
}
