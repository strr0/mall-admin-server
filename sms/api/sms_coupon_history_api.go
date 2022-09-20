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

func (iApi SmsCouponHistoryApi) List(ctx *gin.Context) {
	couponId := ctx.Query("couponId")
	useStatus := ctx.Query("useStatus")
	orderSn := ctx.Query("orderSn")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(couponId, useStatus, orderSn, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
