package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/cms/api"
	"mall-admin-server/cms/service"
)

func init() {
	routers = append(routers, registerCmsPrefrenceAreaRouter)
}

func registerCmsPrefrenceAreaRouter(e *gin.Engine) {
	iService := service.CmsPrefrenceAreaService{}
	iApi := api.CmsPrefrenceAreaApi{
		Service: iService,
	}
	prefrenceArea := e.Group("/prefrenceArea")
	{
		prefrenceArea.GET("/listAll", iApi.ListAll)
	}
}
