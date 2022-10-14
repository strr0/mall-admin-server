package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

// 商品属性管理
type PmsProductAttributeApi struct {
	Service service.PmsProductAttributeService
}

// @Summary      根据分类查询属性列表或参数列表
// @Description  根据分类查询属性列表或参数列表
// @Tags         商品属性管理
// @Accept       json
// @Produce      json
// @Param        cid   path      string  false  "cid"
// @Param        type   query      string  false  "type"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/list/{cid} [get]
func (iApi PmsProductAttributeApi) List(ctx *gin.Context) {
	cid := ctx.Query("cid")
	type_ := ctx.Query("type")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(cid, type_, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      添加商品属性信息
// @Description  添加商品属性信息
// @Tags         商品属性管理
// @Accept       json
// @Produce      json
// @Param        pmsProductAttribute   query      model.PmsProductAttribute  false  "商品属性"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/create [post]
func (iApi PmsProductAttributeApi) Create(ctx *gin.Context) {
	var pmsProductAttribute model.PmsProductAttribute
	_ = ctx.Bind(&pmsProductAttribute)
	err := iApi.Service.Create(pmsProductAttribute)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      修改商品属性信息
// @Description  修改商品属性信息
// @Tags         商品属性管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        pmsProductAttribute   query      model.PmsProductAttribute  false  "商品属性"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/update/{id} [post]
func (iApi PmsProductAttributeApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pmsProductAttribute model.PmsProductAttribute
	_ = ctx.Bind(&pmsProductAttribute)
	err := iApi.Service.Update(id, pmsProductAttribute)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      查询单个商品属性
// @Description  查询单个商品属性
// @Tags         商品属性管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/{id} [get]
func (iApi PmsProductAttributeApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      批量删除商品属性
// @Description  批量删除商品属性
// @Tags         商品属性管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/delete [post]
func (iApi PmsProductAttributeApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      根据商品分类的id获取商品属性及属性分类
// @Description  根据商品分类的id获取商品属性及属性分类
// @Tags         商品属性管理
// @Accept       json
// @Produce      json
// @Param        productCategoryId   path      string  false  "productCategoryId"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/attrInfo/{productCategoryId} [get]
func (iApi PmsProductAttributeApi) GetAttrInfo(ctx *gin.Context) {
	productCategoryId := ctx.Query("productCategoryId")
	info := iApi.Service.GetProductAttrInfo(productCategoryId)
	ctx.JSON(http.StatusOK, util.Data(info))
}
