package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	cms "mall-admin-server/cms/router"
	"mall-admin-server/config"
	"mall-admin-server/docs"
	oms "mall-admin-server/oms/router"
	pms "mall-admin-server/pms/router"
	sms "mall-admin-server/sms/router"
	ums "mall-admin-server/ums/router"
)

var StartCmd = &cobra.Command{
	Use:   "admin",
	Short: "admin",
	Long:  "admin",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// 数据源
func initDataSource() {
	config.InitDataSource()
}

// casbin权限配置
func initPolicy() {
	config.InitPolicy()
}

// 路由
func initRouter() {
	e := gin.Default()
	cms.InitRouter(e)
	oms.InitRouter(e)
	pms.InitRouter(e)
	sms.InitRouter(e)
	ums.InitRouter(e)
	docs.InitSwagger(e) // 引入swagger
	_ = e.Run()
}

func run() {
	initDataSource()
	initPolicy()
	initRouter()
}
