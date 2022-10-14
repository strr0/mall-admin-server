package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/common/api"
	"mall-admin-server/common/service"
	"mall-admin-server/config"
)

func InitRouter(e *gin.Engine) {
	redisdb := config.GetRedis()
	iService := service.RedisService{RedisDB: redisdb}
	iApi := api.RedisApi{Service: iService}
	redis := e.Group("/redis")
	{
		redis.GET("/get", iApi.Get)
		redis.GET("/set", iApi.Set)
	}
}
