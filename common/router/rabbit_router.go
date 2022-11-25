package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/common/api"
	"mall-admin-server/common/service"
	"mall-admin-server/config"
)

func init() {
	routers = append(routers, registerRabbitRouter)
}

func registerRabbitRouter(e *gin.Engine) {
	conn := config.GetRabbit()
	iService := service.RabbitService{Rabbit: conn}
	iApi := api.RabbitApi{Service: iService}
	rabbit := e.Group("/rabbit")
	{
		rabbit.POST("/send", iApi.Send)
		rabbit.GET("/recv", iApi.Recv)
		rabbit.GET("/initQueue", iApi.InitQueue)
	}
}
