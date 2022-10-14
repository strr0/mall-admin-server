package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/sms/api"
	"mall-admin-server/sms/service"
)

func init() {
	routers = append(routers, registerSmsHomeRecommendSubjectRouter)
}

func registerSmsHomeRecommendSubjectRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.SmsHomeRecommendSubjectService{DB: db}
	iApi := api.SmsHomeRecommendSubjectApi{Service: iService}
	homeRecommendSubject := e.Group("/home/recommendSubject").Use(config.AuthCheckRole("homeSubject"))
	{
		homeRecommendSubject.POST("/create", iApi.Create)
		homeRecommendSubject.POST("/update/sort/:id", iApi.UpdateSort)
		homeRecommendSubject.POST("/delete", iApi.Delete)
		homeRecommendSubject.POST("/update/recommendStatus", iApi.UpdateRecommendStatus)
		homeRecommendSubject.GET("/list", iApi.List)
	}
}
