package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/oms/api"
	"mall-admin-server/oms/service"
)

func init() {
	routers = append(routers, registerOmsOrderSettingRouter)
}

func registerOmsOrderSettingRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.OmsOrderSettingService{DB: db}
	iApi := api.OmsOrderSettingApi{Service: iService}
	orderSetting := e.Group("/orderSetting")
	{
		orderSetting.GET("/:id", iApi.GetItem)
		orderSetting.POST("/update/:id", iApi.Update)
	}
}
