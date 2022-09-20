package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/service"
	"mall-admin-server/util"
	"net/http"
)

// 订单设置管理
type OmsOrderSettingApi struct {
	Service service.OmsOrderSettingService
}

// 获取指定订单设置
func (iApi OmsOrderSettingApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 修改指定订单设置
func (iApi OmsOrderSettingApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var omsOrderSetting model.OmsOrderSetting
	_ = ctx.Bind(&omsOrderSetting)
	err := iApi.Service.Update(id, omsOrderSetting)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
