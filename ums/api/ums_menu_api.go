package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 后台菜单管理
type UmsMenuApi struct {
	Service service.UmsMenuService
}

// @Summary      新增菜单
// @Description  新增菜单
// @Tags         后台菜单管理
// @Accept       json
// @Produce      json
// @Param        umsMenu   query      model.UmsMenu  false  "菜单"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /menu/create [post]
func (iApi UmsMenuApi) Create(ctx *gin.Context) {
	var umsMenu model.UmsMenu
	_ = ctx.Bind(&umsMenu)
	err := iApi.Service.Create(umsMenu)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      修改菜单
// @Description  修改菜单
// @Tags         后台菜单管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        umsMenu   query      model.UmsMenu  false  "菜单"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /menu/update/{id} [post]
func (iApi UmsMenuApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var umsMenu model.UmsMenu
	_ = ctx.Bind(&umsMenu)
	err := iApi.Service.Update(id, umsMenu)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      查看菜单
// @Description  查看菜单
// @Tags         后台菜单管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /menu/{id} [get]
func (iApi UmsMenuApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      删除菜单
// @Description  删除菜单
// @Tags         后台菜单管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /menu/delete/{id} [post]
func (iApi UmsMenuApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除失败"))
}

// @Summary      菜单列表
// @Description  菜单列表
// @Tags         后台菜单管理
// @Accept       json
// @Produce      json
// @Param        parentId   path      string  false  "父id"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /menu/list/{parentId} [get]
func (iApi UmsMenuApi) List(ctx *gin.Context) {
	parentId := ctx.Param("parentId")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(parentId, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      获取树形菜单
// @Description  获取树形菜单
// @Tags         后台菜单管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /menu/treeList [get]
func (iApi UmsMenuApi) TreeList(ctx *gin.Context) {
	list := iApi.Service.TreeList()
	ctx.JSON(http.StatusOK, util.Data(list))
}

// @Summary      修改菜单显示状态
// @Description  修改菜单显示状态
// @Tags         后台菜单管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        hidden   query      string  false  "hidden"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /menu/updateHidden/{id} [post]
func (iApi UmsMenuApi) UpdateHidden(ctx *gin.Context) {
	id := ctx.Param("id")
	hidden := ctx.PostForm("hidden")
	err := iApi.Service.UpdateHidden(id, hidden)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
