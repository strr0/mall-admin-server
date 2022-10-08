package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	cmsq "mall-admin-server/cms/query"
	cmsr "mall-admin-server/cms/router"
	"mall-admin-server/config"
	omsq "mall-admin-server/oms/query"
	omsr "mall-admin-server/oms/router"
	pmsq "mall-admin-server/pms/query"
	pmsr "mall-admin-server/pms/router"
	smsq "mall-admin-server/sms/query"
	smsr "mall-admin-server/sms/router"
	umsr "mall-admin-server/ums/router"
)

var StartCmd = &cobra.Command{
	Use:   "admin",
	Short: "admin",
	Long:  "admin",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func initQuery() {
	db := config.GetDb()
	cmsq.SetDefault(db)
	omsq.SetDefault(db)
	pmsq.SetDefault(db)
	smsq.SetDefault(db)
}

func initRouter() {
	e := gin.Default()
	cmsr.InitRouter(e)
	omsr.InitRouter(e)
	pmsr.InitRouter(e)
	smsr.InitRouter(e)
	umsr.InitRouter(e)
	_ = e.Run()
}

func run() {
	initQuery()
	initRouter()
}
