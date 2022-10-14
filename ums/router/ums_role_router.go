package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsRoleRouter)
}

func registerUmsRoleRouter(e *gin.Engine) {
	db := config.GetDb()
	iService := service.UmsRoleService{DB: db}
	iApi := api.UmsRoleApi{Service: iService}
	role := e.Group("/role").Use(config.AuthCheckRole("role"))
	{
		role.POST("/create", iApi.Create)
		role.POST("/update/:id", iApi.Update)
		role.POST("/delete/:id", iApi.Delete)
		role.GET("/listAll", iApi.ListAll)
		role.GET("/list", iApi.List)
		role.POST("/updateStatus/:id", iApi.UpdateStatus)
		role.GET("/listMenu/:roleId", iApi.ListMenu)
		role.GET("/listResource/:roleId", iApi.ListResource)
		role.POST("/allocMenu", iApi.AllocMenu)
		role.POST("/allocResource", iApi.AllocResource)
	}
}
