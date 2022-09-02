package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

type UmsMenuApi struct {
	Service service.UmsMenuService
}

// 新增菜单
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

// 修改菜单
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

// 查看菜单
func (iApi UmsMenuApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 删除菜单
func (iApi UmsMenuApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除失败"))
}

// 菜单列表
func (iApi UmsMenuApi) List(ctx *gin.Context) {
	parentId := ctx.Param("parentId")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(parentId, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// 获取树形菜单
func (iApi UmsMenuApi) TreeList(ctx *gin.Context) {
	list := iApi.Service.TreeList()
	ctx.JSON(http.StatusOK, util.Data(list))
}

// 修改菜单显示状态
func (iApi UmsMenuApi) UpdateHidden(ctx *gin.Context) {
	id := ctx.Param("id")
	hidden := ctx.Query("hidden")
	err := iApi.Service.UpdateHidden(id, hidden)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
