package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/cms/api"
	"mall-admin-server/cms/service"
	"mall-admin-server/config"
)

func init() {
	routers = append(routers, registerCmsPrefrenceAreaRouter)
}

func registerCmsPrefrenceAreaRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.CmsPrefrenceAreaService{DB: db}
	iApi := api.CmsPrefrenceAreaApi{Service: iService}
	prefrenceArea := e.Group("/prefrenceArea")
	{
		prefrenceArea.GET("/listAll", iApi.ListAll)
	}
}
