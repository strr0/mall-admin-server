package main

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/router"
)

func main() {
	engine := gin.Default()
	router.InitRouter(engine)
	_ = engine.Run()
}
