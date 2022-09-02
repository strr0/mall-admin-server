package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

type UmsResourceApi struct {
	Service service.UmsResourceService
}

// 创建资源
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

// 修改资源
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

// 资源详情
func (iApi UmsResourceApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 删除资源
func (iApi UmsResourceApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 资源列表
func (iApi UmsResourceApi) List(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId")
	nameKeyword := ctx.Query("nameKeyword")
	urlKeyword := ctx.Query("urlKeyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(categoryId, nameKeyword, urlKeyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// 所有资源
func (iApi UmsResourceApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}
