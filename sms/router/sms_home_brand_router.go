package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsHomeBrandRouter)
}

func registerSmsHomeBrandRouter(e *gin.Engine) {
	iService := service.SmsHomeBrandService{}
	iApi := api.SmsHomeBrandApi{
		Service: iService,
	}
	homeBrand := e.Group("/home/brand")
	{
		homeBrand.POST("/create", iApi.Create)
		homeBrand.POST("/update/sort/:id", iApi.UpdateSort)
		homeBrand.POST("/delete", iApi.Delete)
		homeBrand.POST("/update/recommendStatus", iApi.UpdateRecommendStatus)
		homeBrand.GET("/list", iApi.List)
	}
}