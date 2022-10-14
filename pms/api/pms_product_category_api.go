package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/service"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
	"net/http"
)

// 商品分类管理
type PmsProductCategoryApi struct {
	Service service.PmsProductCategoryService
}

// @Summary      添加商品分类
// @Description  添加商品分类
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        pmsProductCategoryDto   query      dto.PmsProductCategoryDto  false  "分类"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/create [post]
func (iApi PmsProductCategoryApi) Create(ctx *gin.Context) {
	var pmsProductCategoryDto dto.PmsProductCategoryDto
	_ = ctx.Bind(&pmsProductCategoryDto)
	err := iApi.Service.Create(pmsProductCategoryDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      修改商品分类
// @Description  修改商品分类
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        pmsProductCategoryDto   query      dto.PmsProductCategoryDto  false  "分类"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/update/{id} [post]
func (iApi PmsProductCategoryApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pmsProductCategoryDto dto.PmsProductCategoryDto
	_ = ctx.Bind(&pmsProductCategoryDto)
	err := iApi.Service.Update(id, pmsProductCategoryDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      分页查询商品分类
// @Description  分页查询商品分类
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        parentId   query      string  false  "parentId"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/list [get]
func (iApi PmsProductCategoryApi) List(ctx *gin.Context) {
	parentId := ctx.Param("parentId")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(parentId, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      根据id获取商品分类
// @Description  根据id获取商品分类
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/{id} [get]
func (iApi PmsProductCategoryApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      删除商品分类
// @Description  删除商品分类
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/delete/{id} [post]
func (iApi PmsProductCategoryApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      修改导航栏显示状态
// @Description  修改导航栏显示状态
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        navStatus   query      string  false  "navStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/update/navStatus [post]
func (iApi PmsProductCategoryApi) UpdateNavStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	navStatus := ctx.PostForm("navStatus")
	err := iApi.Service.UpdateNavStatus(ids, navStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      修改显示状态
// @Description  修改显示状态
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        showStatus   query      string  false  "showStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/update/showStatus [post]
func (iApi PmsProductCategoryApi) UpdateShowStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	showStatus := ctx.PostForm("showStatus")
	err := iApi.Service.UpdateShowStatus(ids, showStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      查询所有一级分类及子分类
// @Description  查询所有一级分类及子分类
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /productCategory/list/withChildren [get]
func (iApi PmsProductCategoryApi) ListWithChildren(ctx *gin.Context) {
	list := iApi.Service.TreeList()
	ctx.JSON(http.StatusOK, util.Data(list))
}
