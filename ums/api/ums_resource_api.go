package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 后台资源管理
type UmsResourceApi struct {
	Service service.UmsResourceService
}

// @Summary      创建资源
// @Description  创建资源
// @Tags         后台资源管理
// @Accept       json
// @Produce      json
// @Param        umsResource   query      model.UmsResource  false  "资源"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resource/create [post]
func (iApi UmsResourceApi) Create(ctx *gin.Context) {
	var umsResource model.UmsResource
	_ = ctx.Bind(&umsResource)
	err := iApi.Service.Create(umsResource)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      修改资源
// @Description  修改资源
// @Tags         后台资源管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        umsResource   query      model.UmsResource  false  "资源"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resource/update/{id} [post]
func (iApi UmsResourceApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var umsResource model.UmsResource
	err := iApi.Service.Update(id, umsResource)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      资源详情
// @Description  资源详情
// @Tags         后台资源管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resource/{id} [get]
func (iApi UmsResourceApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      删除资源
// @Description  删除资源
// @Tags         后台资源管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resource/delete/{id} [post]
func (iApi UmsResourceApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      资源列表
// @Description  资源列表
// @Tags         后台资源管理
// @Accept       json
// @Produce      json
// @Param        categoryId   query      string  false  "categoryId"
// @Param        nameKeyword   query      string  false  "nameKeyword"
// @Param        urlKeyword   query      string  false  "urlKeyword"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resource/list [get]
func (iApi UmsResourceApi) List(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId")
	nameKeyword := ctx.Query("nameKeyword")
	urlKeyword := ctx.Query("urlKeyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(categoryId, nameKeyword, urlKeyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      所有资源
// @Description  所有资源
// @Tags         后台资源管理
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resource/listAll [get]
func (iApi UmsResourceApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}
