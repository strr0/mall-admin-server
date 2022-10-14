package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsHomeNewProductRouter)
}

func registerSmsHomeNewProductRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.SmsHomeNewProductService{DB: db}
	iApi := api.SmsHomeNewProductApi{Service: iService}
	homeNewProduct := e.Group("/home/newProduct").Use(config.AuthCheckRole("homeNew"))
	{
		homeNewProduct.POST("/create", iApi.Create)
		homeNewProduct.POST("/update/sort/:id", iApi.UpdateSort)
		homeNewProduct.POST("/delete", iApi.Delete)
		homeNewProduct.POST("/update/recommendStatus", iApi.UpdateRecommendStatus)
		homeNewProduct.GET("/list", iApi.List)
	}
}
