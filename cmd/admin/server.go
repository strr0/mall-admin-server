package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	cms "mall-admin-server/cms/router"
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

func initRouter() {
	e := gin.Default()
	cms.InitRouter(e)
	oms.InitRouter(e)
	pms.InitRouter(e)
	sms.InitRouter(e)
	ums.InitRouter(e)
	_ = e.Run()
}

func run() {
	initRouter()
}
