package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsFlashPromotionRouter)
}

func registerSmsFlashPromotionRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.SmsFlashPromotionService{DB: db}
	iApi := api.SmsFlashPromotionApi{Service: iService}
	flashPromotion := e.Group("/flash")
	{
		flashPromotion.POST("/create", iApi.Create)
		flashPromotion.POST("/update/:id", iApi.Update)
		flashPromotion.POST("/delete/:id", iApi.Delete)
		flashPromotion.POST("/update/status/:id", iApi.UpdateStatus)
		flashPromotion.GET("/:id", iApi.GetItem)
		flashPromotion.GET("/list", iApi.List)
	}
}
