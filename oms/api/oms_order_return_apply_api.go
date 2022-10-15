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

// @Summary      分页查询退货申请
// @Description  分页查询退货申请
// @Tags         订单退货申请管理
// @Accept       json
// @Produce      json
// @Param        queryDto   query      dto.OmsOrderReturnApplyQueryDto  false  "queryDto"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnApply/list [get]
func (iApi OmsOrderReturnApplyApi) List(ctx *gin.Context) {
	var queryDto dto.OmsOrderReturnApplyQueryDto
	_ = ctx.Bind(&queryDto)
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(queryDto, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}

// @Summary      批量删除退货申请
// @Description  批量删除退货申请
// @Tags         订单退货申请管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnApply/delete [post]
func (iApi OmsOrderReturnApplyApi) Delete(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      获取退货申请详情
// @Description  获取退货申请详情
// @Tags         订单退货申请管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnApply/{id} [get]
func (iApi OmsOrderReturnApplyApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      修改退货申请状态
// @Description  修改退货申请状态
// @Tags         订单退货申请管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        statusDto   query      dto.OmsUpdateStatusDto  false  "statusDto"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /returnApply/update/status/{id} [post]
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
