package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/api"
	"mall-admin-server/pms/service"
)

func init() {
	routers = append(routers, registerPmsProductCategoryRouter)
}

func registerPmsProductCategoryRouter(e *gin.Engine) {
	iService := service.PmsProductCategoryService{}
	iApi := api.PmsProductCategoryApi{
		Service: iService,
	}
	category := e.Group("/productCategory")
	{
		category.POST("/create", iApi.Create)
		category.POST("/update/:id", iApi.Update)
		category.GET("/list/:parentId", iApi.List)
		category.GET("/:id", iApi.GetItem)
		category.POST("/delete/:id", iApi.Delete)
		category.POST("/update/navStatus", iApi.UpdateNavStatus)
		category.POST("/update/showStatus", iApi.UpdateShowStatus)
		category.GET("/list/withChildren", iApi.ListWithChildren)
	}
}
