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

// 根据分类查询属性列表或参数列表
func (iApi PmsProductAttributeApi) List(ctx *gin.Context) {
	cid := ctx.Query("cid")
	type_ := ctx.Query("type")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(cid, type_, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// 添加商品属性信息
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

// 修改商品属性信息
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

// 查询单个商品属性
func (iApi PmsProductAttributeApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 批量删除商品属性
func (iApi PmsProductAttributeApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 根据商品分类的id获取商品属性及属性分类
func (iApi PmsProductAttributeApi) GetAttrInfo(ctx *gin.Context) {
	productCategoryId := ctx.Query("productCategoryId")
	info := iApi.Service.GetProductAttrInfo(productCategoryId)
	ctx.JSON(http.StatusOK, util.Data(info))
}
