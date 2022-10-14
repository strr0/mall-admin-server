package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/service"
	"mall-admin-server/sms/service/dto"
	"mall-admin-server/util"
	"net/http"
)

// 优惠券管理
type SmsCouponApi struct {
	Service service.SmsCouponService
}

// @Summary      添加优惠券
// @Description  添加优惠券
// @Tags         优惠券管理
// @Accept       json
// @Produce      json
// @Param        smsCouponDto   query      dto.SmsCouponDto  false  "优惠券"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /coupon/create [post]
func (iApi SmsCouponApi) Create(ctx *gin.Context) {
	var smsCouponDto dto.SmsCouponDto
	_ = ctx.Bind(&smsCouponDto)
	err := iApi.Service.Create(smsCouponDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      删除优惠券
// @Description  删除优惠券
// @Tags         优惠券管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /coupon/delete/{id} [post]
func (iApi SmsCouponApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      修改优惠券
// @Description  修改优惠券
// @Tags         优惠券管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        smsCouponDto   query      dto.SmsCouponDto  false  "优惠券"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /coupon/update/{id} [post]
func (iApi SmsCouponApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var smsCouponDto dto.SmsCouponDto
	_ = ctx.Bind(&smsCouponDto)
	err := iApi.Service.Update(id, smsCouponDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      根据优惠券名称和类型分页获取优惠券列表
// @Description  根据优惠券名称和类型分页获取优惠券列表
// @Tags         优惠券管理
// @Accept       json
// @Produce      json
// @Param        name   query      string  false  "名称"
// @Param        type   query      string  false  "类型"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /coupon/list [get]
func (iApi SmsCouponApi) List(ctx *gin.Context) {
	name := ctx.Query("name")
	type_ := ctx.Query("type")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(name, type_, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      获取单个优惠券的详细信息
// @Description  获取单个优惠券的详细信息
// @Tags         优惠券管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /coupon/{id} [get]
func (iApi SmsCouponApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}
