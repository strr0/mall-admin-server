package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/auth"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 后台用户管理
type UmsAdminApi struct {
	UmsAdminService service.UmsAdminService
	UmsMenuService  service.UmsMenuService
	UmsRoleService  service.UmsRoleService
}

// 注册
func (iApi UmsAdminApi) Register(ctx *gin.Context) {
	var umsAdmin model.UmsAdmin
	_ = ctx.Bind(&umsAdmin)
	err := iApi.UmsAdminService.Register(umsAdmin)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("注册失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("注册成功"))
}

// 登录
func (iApi UmsAdminApi) Login(ctx *gin.Context) {
	var login model.UmsAdmin
	_ = ctx.Bind(&login)
	umsAdmin := iApi.UmsAdminService.GetAdminByUsername(login.Username)
	if umsAdmin == nil {
		ctx.JSON(http.StatusOK, util.Failed("用户不存在"))
		return
	}
	err := auth.CheckPassword(umsAdmin.Password, login.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("密码错误"))
		return
	}
	tokenMap := auth.GenerateToken(umsAdmin.ID, umsAdmin.Username)
	ctx.JSON(http.StatusOK, util.Data(tokenMap))
}

// 刷新token
func (UmsAdminApi) RefreshToken(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	tokenMap := auth.RefreshToken(tokenString)
	ctx.JSON(http.StatusOK, util.Data(tokenMap))
}

// 获取用户信息
func (iApi UmsAdminApi) GetAdminInfo(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	id, name := auth.ParseToken(tokenString)
	info := map[string]interface{}{
		"username": name,
		"menus":    iApi.UmsMenuService.GetMenuList(id),
		"roles":    iApi.UmsRoleService.GetRoleNameList(id),
	}
	ctx.JSON(http.StatusOK, util.Data(info))
}

// 登出
func (UmsAdminApi) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, util.Success("登出成功"))
}

// 列表
func (iApi UmsAdminApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.UmsAdminService.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// 根据id获取用户信息
func (iApi UmsAdminApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.UmsAdminService.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 修改用户信息
func (iApi UmsAdminApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var umsAdmin model.UmsAdmin
	_ = ctx.Bind(&umsAdmin)
	err := iApi.UmsAdminService.Update(id, umsAdmin)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 删除用户
func (iApi UmsAdminApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.UmsAdminService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 更新状态
func (iApi UmsAdminApi) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.PostForm("status")
	err := iApi.UmsAdminService.UpdateStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 角色分配
func (iApi UmsAdminApi) UpdateRole(ctx *gin.Context) {
	adminId := ctx.PostForm("adminId")
	roleIds := ctx.PostFormArray("roleIds")
	err := iApi.UmsAdminService.UpdateRole(adminId, roleIds)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 获取用户角色
func (iApi UmsAdminApi) GetRoleList(ctx *gin.Context) {
	adminId := ctx.Param("adminId")
	list := iApi.UmsRoleService.GetRoleList(adminId)
	ctx.JSON(http.StatusOK, util.Data(list))
}
