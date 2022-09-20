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

func (iApi SmsCouponApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

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

func (iApi SmsCouponApi) List(ctx *gin.Context) {
	name := ctx.Query("name")
	type_ := ctx.Query("type")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(name, type_, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

func (iApi SmsCouponApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}
