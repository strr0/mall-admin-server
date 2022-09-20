package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/service"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
	"net/http"
)

type PmsProductApi struct {
	Service service.PmsProductService
}

func (iApi PmsProductApi) Create(ctx *gin.Context) {
	var pmsProductDto dto.PmsProductDto
	_ = ctx.Bind(&pmsProductDto)
	err := iApi.Service.Create(pmsProductDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

func (iApi PmsProductApi) GetUpdateInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	info := iApi.Service.GetUpdateInfo(id)
	ctx.JSON(http.StatusOK, util.Data(info))
}

func (iApi PmsProductApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pmsProductDto dto.PmsProductDto
	_ = ctx.Bind(&pmsProductDto)
	err := iApi.Service.Update(id, pmsProductDto)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductApi) List(ctx *gin.Context) {
	var pmsProductQueryDto dto.PmsProductQueryDto
	_ = ctx.Bind(&pmsProductQueryDto)
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(pmsProductQueryDto, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

func (iApi PmsProductApi) SimpleList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	list := iApi.Service.SimpleList(keyword)
	ctx.JSON(http.StatusOK, util.Data(list))
}

func (iApi PmsProductApi) UpdateVerifyStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	verifyStatus := ctx.PostForm("verifyStatus")
	detail := ctx.PostForm("detail")
	err := iApi.Service.UpdateVerifyStatus(ids, verifyStatus, detail)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductApi) UpdatePublishStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	publishStatus := ctx.PostForm("publishStatus")
	err := iApi.Service.UpdatePublishStatus(ids, publishStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductApi) UpdateRecommendStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	recommendStatus := ctx.PostForm("recommendStatus")
	err := iApi.Service.UpdateRecommendStatus(ids, recommendStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductApi) UpdateNewStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	newStatus := ctx.PostForm("newStatus")
	err := iApi.Service.UpdateNewStatus(ids, newStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

func (iApi PmsProductApi) UpdateDeleteStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	deleteStatus := ctx.PostForm("newStatus")
	err := iApi.Service.UpdateDeleteStatus(ids, deleteStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
