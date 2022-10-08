package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsResourceCategoryRouter)
}

func registerUmsResourceCategoryRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.UmsResourceCategoryService{DB: db}
	iApi := api.UmsResourceCategoryApi{Service: iService}
	resourceCategory := e.Group("/resourceCategory")
	{
		resourceCategory.GET("/listAll", iApi.ListAll)
		resourceCategory.POST("/create", iApi.Create)
		resourceCategory.POST("/update/:id", iApi.Update)
		resourceCategory.POST("/delete/:id", iApi.Delete)
	}
}
