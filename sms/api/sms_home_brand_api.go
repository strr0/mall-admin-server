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

// @Summary      添加首页推荐品牌
// @Description  添加首页推荐品牌
// @Tags         首页品牌管理
// @Accept       json
// @Produce      json
// @Param        smsHomeBrand   query      model.SmsHomeBrand  false  "品牌"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/brand/create [post]
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

// @Summary      修改推荐品牌排序
// @Description  修改推荐品牌排序
// @Tags         首页品牌管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        sort   query      string  false  "sort"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/brand/update/sort/{id} [post]
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

// @Summary      批量删除推荐品牌
// @Description  批量删除推荐品牌
// @Tags         首页品牌管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/brand/delete [post]
func (iApi SmsHomeBrandApi) Delete(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      批量修改推荐品牌状态
// @Description  批量修改推荐品牌状态
// @Tags         首页品牌管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        recommendStatus   query      string  false  "recommendStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/brand/update/recommendStatus [post]
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

// @Summary      分页查询推荐品牌
// @Description  分页查询推荐品牌
// @Tags         首页品牌管理
// @Accept       json
// @Produce      json
// @Param        brandName   query      string  false  "brandName"
// @Param        recommendStatus   query      string  false  "recommendStatus"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/brand/list [get]
func (iApi SmsHomeBrandApi) List(ctx *gin.Context) {
	brandName := ctx.Query("brandName")
	recommendStatus := ctx.Query("recommendStatus")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(brandName, recommendStatus, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
