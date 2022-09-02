package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

type UmsRoleApi struct {
	Service service.UmsRoleService
}

// 创建角色
func (iApi UmsRoleApi) Create(ctx *gin.Context) {
	var umsRole model.UmsRole
	_ = ctx.Bind(&umsRole)
	err := iApi.Service.Create(umsRole)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 修改角色
func (iApi UmsRoleApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var umsRole model.UmsRole
	err := iApi.Service.Update(id, umsRole)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 删除角色
func (iApi UmsRoleApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 获取所有角色
func (iApi UmsRoleApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// 角色列表
func (iApi UmsRoleApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// 修改角色状态
func (iApi UmsRoleApi) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.Query("status")
	err := iApi.Service.UpdateStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 根据角色id获取菜单
func (iApi UmsRoleApi) ListMenu(ctx *gin.Context) {
	roleId := ctx.Param("roleId")
	menu := iApi.Service.ListMenu(roleId)
	ctx.JSON(http.StatusOK, util.Data(menu))
}

// 根据角色id获取资源
func (iApi UmsRoleApi) ListResource(ctx *gin.Context) {
	roleId := ctx.Param("roleId")
	resource := iApi.Service.ListResource(roleId)
	ctx.JSON(http.StatusOK, util.Data(resource))
}

// 菜单分配
func (iApi UmsRoleApi) AllocMenu(ctx *gin.Context) {
	roleId := ctx.Query("roleId")
	menuIds := ctx.QueryArray("menuIds")
	err := iApi.Service.AllocMenu(roleId, menuIds)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 资源分配
func (iApi UmsRoleApi) AllocResource(ctx *gin.Context) {
	roleId := ctx.Query("roleId")
	resourceIds := ctx.QueryArray("resourceIds")
	err := iApi.Service.AllocResource(roleId, resourceIds)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
