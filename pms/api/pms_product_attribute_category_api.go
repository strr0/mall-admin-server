package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

type PmsProductAttributeCategoryApi struct {
	Service service.PmsProductAttributeCategoryService
}

func (iApi PmsProductAttributeCategoryApi) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := iApi.Service.Create(name)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

func (iApi PmsProductAttributeCategoryApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("name")
	err := iApi.Service.Update(id, name)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductAttributeCategoryApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

func (iApi PmsProductAttributeCategoryApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

func (iApi PmsProductAttributeCategoryApi) List(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

func (iApi PmsProductAttributeCategoryApi) ListWithAttr(ctx *gin.Context) {
	attr := iApi.Service.GetListWithAttr()
	ctx.JSON(http.StatusOK, util.Data(attr))
}
