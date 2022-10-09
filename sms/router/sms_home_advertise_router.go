package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsHomeAdvertiseRouter)
}

func registerSmsHomeAdvertiseRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.SmsHomeAdvertiseService{DB: db}
	iApi := api.SmsHomeAdvertiseApi{Service: iService}
	homeAdvertise := e.Group("/home/advertise")
	{
		homeAdvertise.POST("/create", iApi.Create)
		homeAdvertise.POST("/delete", iApi.Delete)
		homeAdvertise.POST("/update/status/:id", iApi.UpdateStatus)
		homeAdvertise.GET("/:id", iApi.GetItem)
		homeAdvertise.POST("/update/:id", iApi.Update)
		homeAdvertise.GET("/list", iApi.List)
	}
}
