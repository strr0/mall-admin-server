package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/cms/api"
	"mall-admin-server/cms/service"
	"mall-admin-server/config"
)

func init() {
	routers = append(routers, registerCmsSubjectRouter)
}

func registerCmsSubjectRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.CmsSubjectService{DB: db}
	iApi := api.CmsSubjectApi{Service: iService}
	subject := e.Group("/subject")
	{
		subject.GET("/listAll", iApi.ListAll)
		subject.GET("/list", iApi.List)
	}
}
