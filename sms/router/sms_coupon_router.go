package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsCouponRouter)
}

func registerSmsCouponRouter(e *gin.Engine) {
	iService := service.SmsCouponService{}
	iApi := api.SmsCouponApi{
		Service: iService,
	}
	coupon := e.Group("/coupon")
	{
		coupon.POST("/create", iApi.Create)
		coupon.POST("/delete/:id", iApi.Delete)
		coupon.POST("/update/:id", iApi.Update)
		coupon.GET("/list", iApi.List)
		coupon.GET("/:id", iApi.GetItem)
	}
}
