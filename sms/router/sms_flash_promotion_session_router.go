package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsFlashPromotionSessionRouter)
}

func registerSmsFlashPromotionSessionRouter(e *gin.Engine) {
	iService := service.SmsFlashPromotionSessionService{}
	iApi := api.SmsFlashPromotionSessionApi{
		Service: iService,
	}
	flashSession := e.Group("/flashSession")
	{
		flashSession.POST("/create", iApi.Create)
		flashSession.POST("/update/:id", iApi.Update)
		flashSession.POST("/update/status/:id", iApi.UpdateStatus)
		flashSession.POST("/delete/:id", iApi.Delete)
		flashSession.GET("/:id", iApi.GetItem)
		flashSession.GET("/list", iApi.ListAll)
		flashSession.GET("/selectList", iApi.SelectList)
	}
}
