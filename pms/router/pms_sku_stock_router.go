package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/pms/api"
	"mall-admin-server/pms/service"
)

func init() {
	routers = append(routers, registerPmsSkuStockRouter)
}

func registerPmsSkuStockRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.PmsSkuStockService{DB: db}
	iApi := api.PmsSkuStockApi{Service: iService}
	sku := e.Group("/sku").Use(config.AuthCheckRole("product"))
	{
		sku.GET("/:pid", iApi.GetList)
		sku.POST("/update", iApi.Update)
	}
}
