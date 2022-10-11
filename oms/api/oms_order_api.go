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

// @Summary      查询订单
// @Description  查询订单
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/list [get]
func (iApi OmsOrderApi) List(ctx *gin.Context) {
	var queryDto dto.OmsOrderQueryDto
	_ = ctx.Bind(&queryDto)
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(queryDto, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}

// @Summary      批量发货
// @Description  批量发货
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        deliveryParamList   body      []dto.OmsOrderDeliveryDto  false  "deliveryParamList"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/update/delivery [post]
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

// @Summary      批量关闭订单
// @Description  批量关闭订单
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        note   query      string  false  "note"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/update/close [post]
func (iApi OmsOrderApi) Close(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	note := ctx.PostForm("note")
	err := iApi.Service.Close(ids, note)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("操作失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("操作成功"))
}

// @Summary      批量删除订单
// @Description  批量删除订单
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/delete [post]
func (iApi OmsOrderApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      订单详情
// @Description  订单详情
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/{id} [get]
func (iApi OmsOrderApi) Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	detail := iApi.Service.GetDetail(id)
	ctx.JSON(http.StatusOK, util.Data(detail))
}

// @Summary      修改收货人信息
// @Description  修改收货人信息
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        infoDto   query      dto.OmsReceiverInfoDto  false  "infoDto"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/update/receiverInfo [post]
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

// @Summary      修改订单费用信息
// @Description  修改订单费用信息
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        infoDto   query      dto.OmsMoneyInfoDto  false  "infoDto"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/update/moneyInfo [post]
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

// @Summary      修改订单费用信息
// @Description  修改订单费用信息
// @Tags         订单管理
// @Accept       json
// @Produce      json
// @Param        id   query      string  false  "id"
// @Param        note   query      string  false  "note"
// @Param        status   query      string  false  "status"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /order/update/note [post]
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
