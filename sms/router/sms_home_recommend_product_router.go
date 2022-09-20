package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsHomeRecommendProductRouter)
}

func registerSmsHomeRecommendProductRouter(e *gin.Engine) {
	iService := service.SmsHomeRecommendProductService{}
	iApi := api.SmsHomeRecommendProductApi{
		Service: iService,
	}
	homeRecommendProduct := e.Group("/home/recommendProduct")
	{
		homeRecommendProduct.POST("/create", iApi.Create)
		homeRecommendProduct.POST("/update/sort/:id", iApi.UpdateSort)
		homeRecommendProduct.POST("/delete", iApi.Delete)
		homeRecommendProduct.POST("/update/recommendStatus", iApi.UpdateRecommendStatus)
		homeRecommendProduct.GET("/list", iApi.List)
	}
}
