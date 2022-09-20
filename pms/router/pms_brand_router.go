package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/api"
	"mall-admin-server/pms/service"
)

func init() {
	routers = append(routers, registerPmsBrandRouter)
}

func registerPmsBrandRouter(e *gin.Engine) {
	iService := service.PmsBrandService{}
	iApi := api.PmsBrandApi{
		Service: iService,
	}
	brand := e.Group("/brand")
	{
		brand.GET("/listAll", iApi.ListAll)
		brand.POST("/create", iApi.Create)
		brand.POST("/update/:id", iApi.Update)
		brand.POST("/delete/:id", iApi.Delete)
		brand.GET("/list", iApi.List)
		brand.GET("/:id", iApi.GetItem)
		brand.POST("/delete/batch", iApi.DeleteBrand)
		brand.POST("/update/showStatus", iApi.UpdateShowStatus)
		brand.POST("/update/factoryStatus", iApi.UpdateFactoryStatus)
	}
}
