package router

import (
	"github.com/gin-gonic/gin"
	redis "mall-admin-server/common/service"
	"mall-admin-server/config"
	"mall-admin-server/ums/api"
	"mall-admin-server/ums/service"
)

func init() {
	routers = append(routers, registerUmsMemberRouter)
}

func registerUmsMemberRouter(e *gin.Engine) {
	redisService := redis.RedisService{RedisDB: config.GetRedis()}
	iService := service.UmsMemberService{RedisService: redisService}
	iApi := api.UmsMemberApi{Service: iService}
	member := e.Group("/member")
	{
		member.GET("/generateAuthCode", iApi.GenerateAuthCode)
		member.GET("/verifyAuthCode", iApi.VerifyAuthCode)
	}
}
