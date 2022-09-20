package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/service"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
	"net/http"
)

type PmsProductCategoryApi struct {
	Service service.PmsProductCategoryService
}

func (iApi PmsProductCategoryApi) Create(ctx *gin.Context) {
	var pmsProductCategoryDto dto.PmsProductCategoryDto
	_ = ctx.Bind(&pmsProductCategoryDto)
	err := iApi.Service.Create(pmsProductCategoryDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

func (iApi PmsProductCategoryApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pmsProductCategoryDto dto.PmsProductCategoryDto
	_ = ctx.Bind(&pmsProductCategoryDto)
	err := iApi.Service.Update(id, pmsProductCategoryDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductCategoryApi) List(ctx *gin.Context) {
	parentId := ctx.Param("parentId")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(parentId, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

func (iApi PmsProductCategoryApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

func (iApi PmsProductCategoryApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

func (iApi PmsProductCategoryApi) UpdateNavStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	navStatus := ctx.PostForm("navStatus")
	err := iApi.Service.UpdateNavStatus(ids, navStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductCategoryApi) UpdateShowStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	showStatus := ctx.PostForm("showStatus")
	err := iApi.Service.UpdateShowStatus(ids, showStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductCategoryApi) ListWithChildren(ctx *gin.Context) {
	list := iApi.Service.TreeList()
	ctx.JSON(http.StatusOK, util.Data(list))
}
