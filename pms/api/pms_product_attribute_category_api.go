package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

// 商品属性分类管理
type PmsProductAttributeCategoryApi struct {
	Service service.PmsProductAttributeCategoryService
}

// 添加商品属性分类
func (iApi PmsProductAttributeCategoryApi) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := iApi.Service.Create(name)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 修改商品属性分类
func (iApi PmsProductAttributeCategoryApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("name")
	err := iApi.Service.Update(id, name)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 删除单个商品属性分类
func (iApi PmsProductAttributeCategoryApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 获取单个商品属性分类信息
func (iApi PmsProductAttributeCategoryApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 分页获取所有商品属性分类
func (iApi PmsProductAttributeCategoryApi) List(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// 获取所有商品属性分类及其下属性
func (iApi PmsProductAttributeCategoryApi) ListWithAttr(ctx *gin.Context) {
	attr := iApi.Service.GetListWithAttr()
	ctx.JSON(http.StatusOK, util.Data(attr))
}
