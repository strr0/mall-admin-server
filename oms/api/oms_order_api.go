package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/oms/service"
	"mall-admin-server/oms/service/dto"
	"mall-admin-server/util"
	"net/http"
)

// 订单管理
type OmsOrderApi struct {
	Service service.OmsOrderService
}

func (iApi OmsOrderApi) List(ctx *gin.Context) {
	var queryDto dto.OmsOrderQueryDto
	_ = ctx.Bind(&queryDto)
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(queryDto, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}

// 批量发货
func (iApi OmsOrderApi) Delivery(ctx *gin.Context) {
	deliveryParamList := make([]dto.OmsOrderDeliveryDto, 0)
	_ = ctx.Bind(&deliveryParamList)
	err := iApi.Service.Delivery(deliveryParamList)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("操作失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("操作成功"))
}

// 批量关闭订单
func (iApi OmsOrderApi) Close(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	note := ctx.PostForm("note")
	err := iApi.Service.Clone(ids, note)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("操作失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("操作成功"))
}

// 批量删除订单
func (iApi OmsOrderApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 订单详情
func (iApi OmsOrderApi) Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	detail := iApi.Service.GetDetail(id)
	ctx.JSON(http.StatusOK, util.Data(detail))
}

// 修改收货人信息
func (iApi OmsOrderApi) UpdateReceiverInfo(ctx *gin.Context) {
	var infoDto dto.OmsReceiverInfoDto
	_ = ctx.Bind(&infoDto)
	err := iApi.Service.UpdateReceiverInfo(infoDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 修改订单费用信息
func (iApi OmsOrderApi) UpdateMoneyInfo(ctx *gin.Context) {
	var infoDto dto.OmsMoneyInfoDto
	_ = ctx.Bind(&infoDto)
	err := iApi.Service.UpdateMoneyInfo(infoDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 修改订单费用信息
func (iApi OmsOrderApi) UpdateNote(ctx *gin.Context) {
	id := ctx.Query("id")
	note := ctx.Query("note")
	status := ctx.Query("status")
	err := iApi.Service.UpdateNote(id, note, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
