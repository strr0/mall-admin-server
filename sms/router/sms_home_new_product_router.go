package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsHomeNewProductRouter)
}

func registerSmsHomeNewProductRouter(e *gin.Engine) {
	iService := service.SmsHomeNewProductService{}
	iApi := api.SmsHomeNewProductApi{
		Service: iService,
	}
	homeNewProduct := e.Group("/home/newProduct")
	{
		homeNewProduct.POST("/create", iApi.Create)
		homeNewProduct.POST("/update/sort/:id", iApi.UpdateSort)
		homeNewProduct.POST("/delete", iApi.Delete)
		homeNewProduct.POST("/update/recommendStatus", iApi.UpdateRecommendStatus)
		homeNewProduct.GET("/list", iApi.List)
	}
}
