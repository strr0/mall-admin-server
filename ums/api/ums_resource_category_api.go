package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 后台资源分类管理
type UmsResourceCategoryApi struct {
	Service service.UmsResourceCategoryService
}

// 资源分类列表
func (iApi UmsResourceCategoryApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// 创建资源分类
func (iApi UmsResourceCategoryApi) Create(ctx *gin.Context) {
	var umsResourceCategory model.UmsResourceCategory
	_ = ctx.Bind(&umsResourceCategory)
	err := iApi.Service.Create(umsResourceCategory)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 修改资源分类
func (iApi UmsResourceCategoryApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var umsResourceCategory model.UmsResourceCategory
	_ = ctx.Bind(&umsResourceCategory)
	err := iApi.Service.Update(id, umsResourceCategory)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 删除资源分类
func (iApi UmsResourceCategoryApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}
