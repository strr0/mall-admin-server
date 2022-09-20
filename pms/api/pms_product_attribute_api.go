package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

type PmsProductAttributeApi struct {
	Service service.PmsProductAttributeService
}

func (iApi PmsProductAttributeApi) List(ctx *gin.Context) {
	cid := ctx.Query("cid")
	type_ := ctx.Query("type")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(cid, type_, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

func (iApi PmsProductAttributeApi) Create(ctx *gin.Context) {
	var pmsProductAttribute model.PmsProductAttribute
	_ = ctx.Bind(&pmsProductAttribute)
	err := iApi.Service.Create(pmsProductAttribute)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

func (iApi PmsProductAttributeApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pmsProductAttribute model.PmsProductAttribute
	_ = ctx.Bind(&pmsProductAttribute)
	err := iApi.Service.Update(id, pmsProductAttribute)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductAttributeApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

func (iApi PmsProductAttributeApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductAttributeApi) GetAttrInfo(ctx *gin.Context) {
	productCategoryId := ctx.Query("productCategoryId")
	info := iApi.Service.GetProductAttrInfo(productCategoryId)
	ctx.JSON(http.StatusOK, util.Data(info))
}
