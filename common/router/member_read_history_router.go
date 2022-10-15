package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/common/api"
	"mall-admin-server/common/service"
	"mall-admin-server/config"
)

func init() {
	routers = append(routers, registerMemberReadHistoryRouter)
}

func registerMemberReadHistoryRouter(e *gin.Engine) {
	collection := config.GetMongo().Database("mall").Collection("member_read_history")
	iService := service.MemberReadHistoryService{Collection: collection}
	iApi := api.MemberReadHistoryApi{Service: iService}
	history := e.Group("/memberReadHistory")
	{
		history.POST("/create", iApi.Create)
		history.POST("/delete", iApi.Delete)
		history.GET("/list", iApi.List)
	}
}
