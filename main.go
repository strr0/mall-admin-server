package main

import (
	"github.com/gin-gonic/gin"
	_ "mall-admin-server/config"
	"mall-admin-server/ums/router"
)

func main() {
	engine := gin.Default()
	router.InitRouter(engine)
	_ = engine.Run()
}
