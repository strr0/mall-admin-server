package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/config"
)

var routers = make([]func(*gin.Engine), 0)

func InitRouter(e *gin.Engine) {
	e.Use(config.Cors)
	for _, f := range routers {
		f(e)
	}
}
