package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/pms/api"
	"mall-admin-server/pms/service"
)

func init() {
	routers = append(routers, registerPmsProductRouter)
}

func registerPmsProductRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.PmsProductService{DB: db}
	iApi := api.PmsProductApi{Service: iService}
	product := e.Group("/product").Use(config.AuthCheckRole("product"))
	{
		product.POST("/create", iApi.Create)
		product.GET("/updateInfo/:id", iApi.GetUpdateInfo)
		product.POST("/update/:id", iApi.Update)
		product.GET("/list", iApi.List)
		product.GET("/simpleList", iApi.SimpleList)
		product.POST("/update/verifyStatus", iApi.UpdateVerifyStatus)
		product.POST("/update/publishStatus", iApi.UpdatePublishStatus)
		product.POST("/update/recommendStatus", iApi.UpdateRecommendStatus)
		product.POST("/update/newStatus", iApi.UpdateNewStatus)
		product.POST("/update/deleteStatus", iApi.UpdateDeleteStatus)
	}
}
