package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 优惠券领取记录管理
type SmsCouponHistoryApi struct {
	Service service.SmsCouponHistoryService
}

// @Summary      根据优惠券id，使用状态，订单编号分页获取领取记录
// @Description  根据优惠券id，使用状态，订单编号分页获取领取记录
// @Tags         优惠券领取记录管理
// @Accept       json
// @Produce      json
// @Param        couponId   query      string  false  "couponId"
// @Param        useStatus   query      string  false  "useStatus"
// @Param        orderSn   query      string  false  "orderSn"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /couponHistory/list [get]
func (iApi SmsCouponHistoryApi) List(ctx *gin.Context) {
	couponId := ctx.Query("couponId")
	useStatus := ctx.Query("useStatus")
	orderSn := ctx.Query("orderSn")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(couponId, useStatus, orderSn, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
