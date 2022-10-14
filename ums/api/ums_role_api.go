package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 后台用户角色管理
type UmsRoleApi struct {
	Service service.UmsRoleService
}

// @Summary      创建角色
// @Description  创建角色
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        umsRole   query      model.UmsRole  false  "角色"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/create [post]
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

// @Summary      修改角色
// @Description  修改角色
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        umsRole   query      model.UmsRole  false  "角色"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/update/{id} [post]
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

// @Summary      删除角色
// @Description  删除角色
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/delete/{id} [post]
func (iApi UmsRoleApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      获取所有角色
// @Description  获取所有角色
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/listAll [get]
func (iApi UmsRoleApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// @Summary      角色列表
// @Description  角色列表
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        keyword   query      string  false  "名称"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/list [get]
func (iApi UmsRoleApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      修改角色状态
// @Description  修改角色状态
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        status   query      string  false  "status"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/updateStatus/{id} [post]
func (iApi UmsRoleApi) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.PostForm("status")
	err := iApi.Service.UpdateStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      根据角色id获取菜单
// @Description  根据角色id获取菜单
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        roleId   path      string  false  "roleId"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/listMenu/{roleId} [get]
func (iApi UmsRoleApi) ListMenu(ctx *gin.Context) {
	roleId := ctx.Param("roleId")
	menu := iApi.Service.ListMenu(roleId)
	ctx.JSON(http.StatusOK, util.Data(menu))
}

// @Summary      根据角色id获取资源
// @Description  根据角色id获取资源
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        roleId   path      string  false  "roleId"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/listResource/{roleId} [get]
func (iApi UmsRoleApi) ListResource(ctx *gin.Context) {
	roleId := ctx.Param("roleId")
	resource := iApi.Service.ListResource(roleId)
	ctx.JSON(http.StatusOK, util.Data(resource))
}

// @Summary      菜单分配
// @Description  菜单分配
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        roleId   query      string  false  "roleId"
// @Param        menuIds   query      []string  false  "menuIds"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/allocMenu [post]
func (iApi UmsRoleApi) AllocMenu(ctx *gin.Context) {
	roleId := ctx.PostForm("roleId")
	menuIds := ctx.PostFormArray("menuIds")
	err := iApi.Service.AllocMenu(roleId, menuIds)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      资源分配
// @Description  资源分配
// @Tags         后台用户角色管理
// @Accept       json
// @Produce      json
// @Param        roleId   query      string  false  "roleId"
// @Param        resourceIds   query      []string  false  "resourceIds"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /role/allocResource [post]
func (iApi UmsRoleApi) AllocResource(ctx *gin.Context) {
	roleId := ctx.PostForm("roleId")
	resourceIds := ctx.PostFormArray("resourceIds")
	err := iApi.Service.AllocResource(roleId, resourceIds)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
