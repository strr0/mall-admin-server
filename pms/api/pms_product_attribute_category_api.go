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

// @Summary      添加商品属性分类
// @Description  添加商品属性分类
// @Tags         商品属性分类管理
// @Accept       json
// @Produce      json
// @Param        name   query      string  false  "name"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/category/create [post]
func (iApi PmsProductAttributeCategoryApi) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := iApi.Service.Create(name)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      修改商品属性分类
// @Description  修改商品属性分类
// @Tags         商品属性分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        name   query      string  false  "name"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/category/update/{id} [post]
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

// @Summary      删除单个商品属性分类
// @Description  删除单个商品属性分类
// @Tags         商品属性分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/category/delete/{id} [post]
func (iApi PmsProductAttributeCategoryApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      获取单个商品属性分类信息
// @Description  获取单个商品属性分类信息
// @Tags         商品属性分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/category/{id} [get]
func (iApi PmsProductAttributeCategoryApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      分页获取所有商品属性分类
// @Description  分页获取所有商品属性分类
// @Tags         商品属性分类管理
// @Accept       json
// @Produce      json
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/category/list [get]
func (iApi PmsProductAttributeCategoryApi) List(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      获取所有商品属性分类及其下属性
// @Description  获取所有商品属性分类及其下属性
// @Tags         商品属性分类管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productAttribute/category/list/withAttr [get]
func (iApi PmsProductAttributeCategoryApi) ListWithAttr(ctx *gin.Context) {
	attr := iApi.Service.GetListWithAttr()
	ctx.JSON(http.StatusOK, util.Data(attr))
}
