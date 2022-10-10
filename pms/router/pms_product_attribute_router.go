package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/pms/api"
	"mall-admin-server/pms/service"
)

func init() {
	routers = append(routers, registerPmsProductAttributeRouter)
}

func registerPmsProductAttributeRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.PmsProductAttributeService{DB: db}
	iApi := api.PmsProductAttributeApi{Service: iService}
	attribute := e.Group("/productAttribute")
	{
		attribute.GET("/list/:cid", iApi.List)
		attribute.POST("/create", iApi.Create)
		attribute.POST("/update/:id", iApi.Update)
		attribute.GET("/:id", iApi.GetItem)
		attribute.POST("/delete", iApi.Delete)
		attribute.GET("/attrInfo/:productCategoryId", iApi.GetAttrInfo)
	}
}
