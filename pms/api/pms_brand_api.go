package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

type PmsBrandApi struct {
	Service service.PmsBrandService
}

func (iApi PmsBrandApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

func (iApi PmsBrandApi) Create(ctx *gin.Context) {
	var pmsBrand model.PmsBrand
	_ = ctx.Bind(&pmsBrand)
	err := iApi.Service.Create(pmsBrand)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

func (iApi PmsBrandApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pmsBrand model.PmsBrand
	_ = ctx.Bind(&pmsBrand)
	err := iApi.Service.Update(id, pmsBrand)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsBrandApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

func (iApi PmsBrandApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

func (iApi PmsBrandApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

func (iApi PmsBrandApi) DeleteBrand(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.DeleteBrand(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

func (iApi PmsBrandApi) UpdateShowStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	showStatus := ctx.PostForm("showStatus")
	err := iApi.Service.UpdateShowStatus(ids, showStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Success("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsBrandApi) UpdateFactoryStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	factoryStatus := ctx.PostForm("factoryStatus")
	err := iApi.Service.UpdateFactoryStatus(ids, factoryStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Success("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}