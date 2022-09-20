package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/oms/api"
	"mall-admin-server/oms/service"
)

func init() {
	routers = append(routers, registerOmsOrderReturnApplyRouter)
}

func registerOmsOrderReturnApplyRouter(e *gin.Engine) {
	iService := service.OmsOrderReturnApplyService{}
	iApi := api.OmsOrderReturnApplyApi{
		Service: iService,
	}
	returnApply := e.Group("/returnApply")
	{
		returnApply.GET("/list", iApi.List)
		returnApply.POST("/delete", iApi.Delete)
		returnApply.GET("/:id", iApi.GetItem)
		returnApply.POST("/update/status/:id", iApi.UpdateStatus)
	}
}
