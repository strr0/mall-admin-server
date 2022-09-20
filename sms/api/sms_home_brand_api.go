package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 首页品牌管理
type SmsHomeBrandApi struct {
	Service service.SmsHomeBrandService
}

// 添加首页推荐品牌
func (iApi SmsHomeBrandApi) Create(ctx *gin.Context) {
	var smsHomeBrand model.SmsHomeBrand
	_ = ctx.Bind(&smsHomeBrand)
	err := iApi.Service.Create(smsHomeBrand)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 修改推荐品牌排序
func (iApi SmsHomeBrandApi) UpdateSort(ctx *gin.Context) {
	id := ctx.Param("id")
	sort := ctx.PostForm("sort")
	err := iApi.Service.UpdateSort(id, sort)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 批量删除推荐品牌
func (iApi SmsHomeBrandApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 批量修改推荐品牌状态
func (iApi SmsHomeBrandApi) UpdateRecommendStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	recommendStatus := ctx.PostForm("recommendStatus")
	err := iApi.Service.UpdateRecommendStatus(ids, recommendStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("秀嘎成功"))
}

// 分页查询推荐品牌
func (iApi SmsHomeBrandApi) List(ctx *gin.Context) {
	brandName := ctx.Query("brandName")
	recommendStatus := ctx.Query("recommendStatus")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(brandName, recommendStatus, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
