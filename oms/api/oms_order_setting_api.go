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

// @Summary      获取指定订单设置
// @Description  获取指定订单设置
// @Tags         订单设置管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /orderSetting/{id} [get]
func (iApi OmsOrderSettingApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      修改指定订单设置
// @Description  修改指定订单设置
// @Tags         订单设置管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        omsOrderSetting   query      model.OmsOrderSetting  false  "设置"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /orderSetting/update/{id} [get]
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
