package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsAdminRouter)
}

func registerUmsAdminRouter(e *gin.Engine) {
	umsAdminService := service.UmsAdminService{}
	umsMenuService := service.UmsMenuService{}
	umsRoleService := service.UmsRoleService{}
	iApi := api.UmsAdminApi{
		UmsAdminService: umsAdminService,
		UmsMenuService:  umsMenuService,
		UmsRoleService:  umsRoleService,
	}
	admin := e.Group("/admin")
	{
		admin.POST("/register", iApi.Register)
		admin.POST("/login", iApi.Login)
		admin.GET("/refreshToken", iApi.RefreshToken)
		admin.GET("/info", iApi.GetAdminInfo)
		admin.GET("/logout", iApi.Logout)
		admin.GET("/list", iApi.List)
		admin.GET("/:id", iApi.GetItem)
		admin.POST("/update/:id", iApi.Update)
		admin.POST("/delete/:id", iApi.Delete)
		admin.POST("/updateStatus/:id", iApi.UpdateStatus)
		admin.POST("/role/update", iApi.UpdateRole)
		admin.GET("/role/:adminId", iApi.GetRoleList)
	}
}
