package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/service"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
	"net/http"
)

// 商品管理
type PmsProductApi struct {
	Service service.PmsProductService
}

// @Summary      创建商品
// @Description  创建商品
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        pmsProductDto   query      dto.PmsProductDto  false  "商品"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/create [post]
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

// @Summary      根据商品id获取商品编辑信息
// @Description  根据商品id获取商品编辑信息
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/updateInfo/{id} [get]
func (iApi PmsProductApi) GetUpdateInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	info := iApi.Service.GetUpdateInfo(id)
	ctx.JSON(http.StatusOK, util.Data(info))
}

// @Summary      更新商品
// @Description  更新商品
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        pmsProductDto   query      dto.PmsProductDto  false  "商品"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/update/{id} [post]
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

// @Summary      查询商品
// @Description  查询商品
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        pmsProductQueryDto   query      dto.PmsProductQueryDto  false  "参数"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/list [get]
func (iApi PmsProductApi) List(ctx *gin.Context) {
	var pmsProductQueryDto dto.PmsProductQueryDto
	_ = ctx.Bind(&pmsProductQueryDto)
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(pmsProductQueryDto, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      根据商品名称或货号模糊查询
// @Description  根据商品名称或货号模糊查询
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        keyword   query      string  false  "keyword"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/simpleList [get]
func (iApi PmsProductApi) SimpleList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	list := iApi.Service.SimpleList(keyword)
	ctx.JSON(http.StatusOK, util.Data(list))
}

// @Summary      批量修改审核状态
// @Description  批量修改审核状态
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        verifyStatus   query      string  false  "verifyStatus"
// @Param        detail   query      string  false  "detail"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/update/verifyStatus [post]
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

// @Summary      批量上下架商品
// @Description  批量上下架商品
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        publishStatus   query      string  false  "publishStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/update/publishStatus [post]
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

// @Summary      批量推荐商品
// @Description  批量推荐商品
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        recommendStatus   query      string  false  "recommendStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/update/recommendStatus [post]
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

// @Summary      批量设为新品
// @Description  批量设为新品
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        newStatus   query      string  false  "newStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/update/newStatus [post]
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

// @Summary      批量修改删除状态
// @Description  批量修改删除状态
// @Tags         商品管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        deleteStatus   query      string  false  "deleteStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /product/update/deleteStatus [post]
func (iApi PmsProductApi) UpdateDeleteStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	deleteStatus := ctx.PostForm("deleteStatus")
	err := iApi.Service.UpdateDeleteStatus(ids, deleteStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
