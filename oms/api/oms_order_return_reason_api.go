package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/service"
	"mall-admin-server/util"
	"net/http"
)

// 退货原因管理
type OmsOrderReturnReasonApi struct {
	Service service.OmsOrderReturnReasonService
}

func (iApi OmsOrderReturnReasonApi) Create(ctx *gin.Context) {
	var omsOrderReturnReason model.OmsOrderReturnReason
	_ = ctx.Bind(&omsOrderReturnReason)
	err := iApi.Service.Create(omsOrderReturnReason)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存成功"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

func (iApi OmsOrderReturnReasonApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var omsOrderReturnReason model.OmsOrderReturnReason
	_ = ctx.Bind(&omsOrderReturnReason)
	err := iApi.Service.Update(id, omsOrderReturnReason)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改成功"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi OmsOrderReturnReasonApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除成功"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

func (iApi OmsOrderReturnReasonApi) List(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}

func (iApi OmsOrderReturnReasonApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

func (iApi OmsOrderReturnReasonApi) UpdateStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	status := ctx.PostForm("status")
	err := iApi.Service.UpdateStatus(ids, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改成功"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
