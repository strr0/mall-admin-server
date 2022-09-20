package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsFlashPromotionProductRelationRouter)
}

func registerSmsFlashPromotionProductRelationRouter(e *gin.Engine) {
	iService := service.SmsFlashPromotionProductRelationService{}
	iApi := api.SmsFlashPromotionProductRelationApi{
		Service: iService,
	}
	flashProductRelation := e.Group("/flashProductRelation")
	{
		flashProductRelation.POST("/create", iApi.Create)
		flashProductRelation.POST("/update/:id", iApi.Update)
		flashProductRelation.POST("/delete/:id", iApi.Delete)
		flashProductRelation.GET("/:id", iApi.GetItem)
		flashProductRelation.GET("/list", iApi.List)
	}
}
