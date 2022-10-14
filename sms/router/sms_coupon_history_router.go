package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsCouponHistoryRouter)
}

func registerSmsCouponHistoryRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.SmsCouponHistoryService{DB: db}
	iApi := api.SmsCouponHistoryApi{Service: iService}
	couponHistory := e.Group("/couponHistory").Use(config.AuthCheckRole("coupon"))
	{
		couponHistory.GET("/list", iApi.List)
	}
}
