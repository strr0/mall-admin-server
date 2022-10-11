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

// @Summary      注册
// @Description  注册
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        umsAdmin   query      model.UmsAdmin  false  "用户信息"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/register [post]
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

// @Summary      登录
// @Description  登录
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        username   query      string  false  "用户名"
// @Param        password   query      string  false  "密码"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/login [post]
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

// @Summary      刷新token
// @Description  刷新token
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/refreshToken [get]
func (UmsAdminApi) RefreshToken(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	tokenMap := auth.RefreshToken(tokenString)
	ctx.JSON(http.StatusOK, util.Data(tokenMap))
}

// @Summary      获取用户信息
// @Description  获取用户信息
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/info [get]
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

// @Summary      登出
// @Description  登出
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/logout [get]
func (UmsAdminApi) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, util.Success("登出成功"))
}

// @Summary      列表
// @Description  列表
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        keyword   query      string  false  "名称"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/list [get]
func (iApi UmsAdminApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.UmsAdminService.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      根据id获取用户信息
// @Description  根据id获取用户信息
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/{id} [get]
func (iApi UmsAdminApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.UmsAdminService.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      修改用户信息
// @Description  修改用户信息
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        umsAdmin   query      model.UmsAdmin  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/update/{id} [post]
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

// @Summary      删除用户
// @Description  删除用户
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/delete/{id} [post]
func (iApi UmsAdminApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.UmsAdminService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      更新状态
// @Description  更新状态
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        status   query      string  false  "status"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/updateStatus/{id} [post]
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

// @Summary      角色分配
// @Description  角色分配
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        adminId   query      string  false  "adminId"
// @Param        roleIds   query      []string  false  "roleIds"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/role/update [post]
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

// @Summary      获取用户角色
// @Description  获取用户角色
// @Tags         后台用户管理
// @Accept       json
// @Produce      json
// @Param        adminId   path      string  false  "adminId"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /admin/role/{adminId} [get]
func (iApi UmsAdminApi) GetRoleList(ctx *gin.Context) {
	adminId := ctx.Param("adminId")
	list := iApi.UmsRoleService.GetRoleList(adminId)
	ctx.JSON(http.StatusOK, util.Data(list))
}
