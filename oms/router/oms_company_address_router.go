package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/oms/api"
	"mall-admin-server/oms/service"
)

func init() {
	routers = append(routers, registerOmsCompanyAddressRouter)
}

func registerOmsCompanyAddressRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.OmsCompanyAddressService{DB: db}
	iApi := api.OmsCompanyAddressApi{Service: iService}
	companyAddress := e.Group("/companyAddress")
	{
		companyAddress.GET("/list", iApi.List)
	}
}
