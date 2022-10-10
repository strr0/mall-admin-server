package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/oms/api"
	"mall-admin-server/oms/service"
)

func init() {
	routers = append(routers, registerOmsOrderReturnReasonRouter)
}

func registerOmsOrderReturnReasonRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.OmsOrderReturnReasonService{DB: db}
	iApi := api.OmsOrderReturnReasonApi{Service: iService}
	returnReason := e.Group("/returnReason")
	{
		returnReason.POST("/create", iApi.Create)
		returnReason.POST("/update/:id", iApi.Update)
		returnReason.POST("/delete", iApi.Delete)
		returnReason.GET("/list", iApi.List)
		returnReason.GET("/:id", iApi.GetItem)
		returnReason.POST("/update/status", iApi.UpdateStatus)
	}
}
