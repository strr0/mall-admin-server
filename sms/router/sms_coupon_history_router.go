package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsCouponHistoryRouter)
}

func registerSmsCouponHistoryRouter(e *gin.Engine) {
	iService := service.SmsCouponHistoryService{}
	iApi := api.SmsCouponHistoryApi{
		Service: iService,
	}
	couponHistory := e.Group("/couponHistory")
	{
		couponHistory.GET("/list", iApi.List)
	}
}
