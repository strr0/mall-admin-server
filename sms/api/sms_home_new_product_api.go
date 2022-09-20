package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 首页新品管理
type SmsHomeNewProductApi struct {
	Service service.SmsHomeNewProductService
}

// 添加首页新品
func (iApi SmsHomeNewProductApi) Create(ctx *gin.Context) {
	var smsHomeNewProduct model.SmsHomeNewProduct
	_ = ctx.Bind(&smsHomeNewProduct)
	err := iApi.Service.Create(smsHomeNewProduct)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 修改首页新品排序
func (iApi SmsHomeNewProductApi) UpdateSort(ctx *gin.Context) {
	id := ctx.Param("id")
	sort := ctx.PostForm("sort")
	err := iApi.Service.UpdateSort(id, sort)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 批量删除首页新品
func (iApi SmsHomeNewProductApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 批量修改首页新品状态
func (iApi SmsHomeNewProductApi) UpdateRecommendStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	recommendStatus := ctx.PostForm("recommendStatus")
	err := iApi.Service.UpdateRecommendStatus(ids, recommendStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 分页查询首页新品
func (iApi SmsHomeNewProductApi) List(ctx *gin.Context) {
	productName := ctx.Query("productName")
	recommendStatus := ctx.Query("recommendStatus")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(productName, recommendStatus, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
