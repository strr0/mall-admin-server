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

// @Summary      添加首页新品
// @Description  添加首页新品
// @Tags         首页新品管理
// @Accept       json
// @Produce      json
// @Param        smsHomeNewProduct   query      model.SmsHomeNewProduct  false  "新品"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/newProduct/create [post]
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

// @Summary      修改首页新品排序
// @Description  修改首页新品排序
// @Tags         首页新品管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        sort   query      string  false  "sort"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/newProduct/update/sort/{id} [post]
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

// @Summary      批量删除首页新品
// @Description  批量删除首页新品
// @Tags         首页新品管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/newProduct/delete [post]
func (iApi SmsHomeNewProductApi) Delete(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      批量修改首页新品状态
// @Description  批量修改首页新品状态
// @Tags         首页新品管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        recommendStatus   query      string  false  "recommendStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/newProduct/update/recommendStatus [post]
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

// @Summary      分页查询首页新品
// @Description  分页查询首页新品
// @Tags         首页新品管理
// @Accept       json
// @Produce      json
// @Param        productName   query      string  false  "productName"
// @Param        recommendStatus   query      string  false  "recommendStatus"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/newProduct/list [get]
func (iApi SmsHomeNewProductApi) List(ctx *gin.Context) {
	productName := ctx.Query("productName")
	recommendStatus := ctx.Query("recommendStatus")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(productName, recommendStatus, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
