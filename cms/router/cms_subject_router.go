package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/cms/api"
	"mall-admin-server/cms/service"
)

func init() {
	routers = append(routers, registerCmsSubjectRouter)
}

func registerCmsSubjectRouter(e *gin.Engine) {
	iService := service.CmsSubjectService{}
	iApi := api.CmsSubjectApi{
		Service: iService,
	}
	subject := e.Group("/subject")
	{
		subject.GET("/listAll", iApi.ListAll)
		subject.GET("/list", iApi.List)
	}
}
