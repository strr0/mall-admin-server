package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/oms/service"
	"mall-admin-server/oms/service/dto"
	"mall-admin-server/util"
	"net/http"
)

// 订单退货申请管理
type OmsOrderReturnApplyApi struct {
	Service service.OmsOrderReturnApplyService
}

func (iApi OmsOrderReturnApplyApi) List(ctx *gin.Context) {
	var queryDto dto.OmsOrderReturnApplyQueryDto
	_ = ctx.Bind(&queryDto)
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(queryDto, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}

func (iApi OmsOrderReturnApplyApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

func (iApi OmsOrderReturnApplyApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

func (iApi OmsOrderReturnApplyApi) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var statusDto dto.OmsUpdateStatusDto
	err := iApi.Service.UpdateStatus(id, statusDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
