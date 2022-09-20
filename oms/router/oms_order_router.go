package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/oms/api"
	"mall-admin-server/oms/service"
)

func init() {
	routers = append(routers, registerOmsOrderRouter)
}

func registerOmsOrderRouter(e *gin.Engine) {
	iService := service.OmsOrderService{}
	iApi := api.OmsOrderApi{
		Service: iService,
	}
	order := e.Group("/order")
	{
		order.GET("/list", iApi.List)
		order.POST("/update/delivery", iApi.Delivery)
		order.POST("/update/close", iApi.Close)
		order.POST("/delete", iApi.Delete)
		order.GET("/:id", iApi.Detail)
		order.POST("/update/receiverInfo", iApi.UpdateReceiverInfo)
		order.POST("/update/moneyInfo", iApi.UpdateMoneyInfo)
		order.POST("/update/note", iApi.UpdateNote)
	}
}
