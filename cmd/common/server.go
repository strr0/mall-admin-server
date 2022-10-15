package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"mall-admin-server/common/router"
	"mall-admin-server/config"
	"mall-admin-server/docs"
)

var StartCmd = &cobra.Command{
	Use:   "common",
	Short: "common",
	Long:  "common",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// 路由
func initRouter() {
	e := gin.Default()
	router.InitRouter(e)
	docs.InitSwagger(e) // 引入swagger
	_ = e.Run()
}

func run() {
	//config.InitRedis()
	//defer config.CloseRedis()
	config.InitMongo()
	defer config.CLoseMongo()
	initRouter()
}
