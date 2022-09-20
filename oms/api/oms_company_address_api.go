package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/oms/service"
	"mall-admin-server/util"
	"net/http"
)

// 收货地址管理
type OmsCompanyAddressApi struct {
	Service service.OmsCompanyAddressService
}

func (iApi OmsCompanyAddressApi) List(ctx *gin.Context) {
	list := iApi.Service.List()
	ctx.JSON(http.StatusOK, util.Data(list))
}
