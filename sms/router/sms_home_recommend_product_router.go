package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsHomeRecommendProductRouter)
}

func registerSmsHomeRecommendProductRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.SmsHomeRecommendProductService{DB: db}
	iApi := api.SmsHomeRecommendProductApi{Service: iService}
	homeRecommendProduct := e.Group("/home/recommendProduct").Use(config.AuthCheckRole("homeHot"))
	{
		homeRecommendProduct.POST("/create", iApi.Create)
		homeRecommendProduct.POST("/update/sort/:id", iApi.UpdateSort)
		homeRecommendProduct.POST("/delete", iApi.Delete)
		homeRecommendProduct.POST("/update/recommendStatus", iApi.UpdateRecommendStatus)
		homeRecommendProduct.GET("/list", iApi.List)
	}
}
