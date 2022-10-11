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

// @Summary      获取所有收货地址
// @Description  获取所有收货地址
// @Tags         收货地址管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /companyAddress/list [get]
func (iApi OmsCompanyAddressApi) List(ctx *gin.Context) {
	list := iApi.Service.List()
	ctx.JSON(http.StatusOK, util.Data(list))
}
