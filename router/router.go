package router

import (
	"github.com/gin-gonic/gin"
	cms "mall-admin-server/cms/router"
	oms "mall-admin-server/oms/router"
	pms "mall-admin-server/pms/router"
	sms "mall-admin-server/sms/router"
	ums "mall-admin-server/ums/router"
)

func InitRouter(e *gin.Engine) {
	cms.InitRouter(e)
	oms.InitRouter(e)
	pms.InitRouter(e)
	sms.InitRouter(e)
	ums.InitRouter(e)
}
