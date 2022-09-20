package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/api"
	"mall-admin-server/pms/service"
)

func init() {
	routers = append(routers, registerPmsSkuStockRouter)
}

func registerPmsSkuStockRouter(e *gin.Engine) {
	iService := service.PmsSkuStockService{}
	iApi := api.PmsSkuStockApi{
		Service: iService,
	}
	sku := e.Group("/sku")
	{
		sku.GET("/:pid", iApi.GetList)
		sku.POST("/update", iApi.Update)
	}
}
