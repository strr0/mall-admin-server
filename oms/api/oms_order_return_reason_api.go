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

// @Summary      添加退货原因
// @Description  添加退货原因
// @Tags         退货原因管理
// @Accept       json
// @Produce      json
// @Param        omsOrderReturnReason   query      model.OmsOrderReturnReason  false  "原因"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnReason/create [post]
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

// @Summary      修改退货原因
// @Description  修改退货原因
// @Tags         退货原因管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        omsOrderReturnReason   query      model.OmsOrderReturnReason  false  "原因"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnReason/update/{id} [post]
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

// @Summary      批量删除退货原因
// @Description  批量删除退货原因
// @Tags         退货原因管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        omsOrderReturnReason   query      model.OmsOrderReturnReason  false  "原因"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnReason/delete [post]
func (iApi OmsOrderReturnReasonApi) Delete(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除成功"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      分页查询退货原因
// @Description  分页查询退货原因
// @Tags         退货原因管理
// @Accept       json
// @Produce      json
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnReason/list [get]
func (iApi OmsOrderReturnReasonApi) List(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}

// @Summary      获取单个退货原因详情信息
// @Description  获取单个退货原因详情信息
// @Tags         退货原因管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnReason/{id} [get]
func (iApi OmsOrderReturnReasonApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      修改退货原因启用状态
// @Description  修改退货原因启用状态
// @Tags         退货原因管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        status   query      string  false  "status"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnReason/update/status [post]
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
