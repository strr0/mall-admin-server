package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 限时购和商品关系管理
type SmsFlashPromotionProductRelationApi struct {
	Service service.SmsFlashPromotionProductRelationService
}

// @Summary      批量选择商品添加关联
// @Description  批量选择商品添加关联
// @Tags         限时购和商品关系管理
// @Accept       json
// @Produce      json
// @Param        smsFlashPromotionProductRelation   query      model.SmsFlashPromotionProductRelation  false  "关联"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashProductRelation/create [post]
func (iApi SmsFlashPromotionProductRelationApi) Create(ctx *gin.Context) {
	var smsFlashPromotionProductRelation model.SmsFlashPromotionProductRelation
	_ = ctx.Bind(&smsFlashPromotionProductRelation)
	err := iApi.Service.Create(smsFlashPromotionProductRelation)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      修改关联信息
// @Description  修改关联信息
// @Tags         限时购和商品关系管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        smsFlashPromotionProductRelation   query      model.SmsFlashPromotionProductRelation  false  "关联"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashProductRelation/update/{id} [post]
func (iApi SmsFlashPromotionProductRelationApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var smsFlashPromotionProductRelation model.SmsFlashPromotionProductRelation
	_ = ctx.Bind(&smsFlashPromotionProductRelation)
	err := iApi.Service.Update(id, smsFlashPromotionProductRelation)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      删除关联
// @Description  删除关联
// @Tags         限时购和商品关系管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashProductRelation/delete/{id} [post]
func (iApi SmsFlashPromotionProductRelationApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      获取管理商品促销信息
// @Description  获取管理商品促销信息
// @Tags         限时购和商品关系管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashProductRelation/{id} [get]
func (iApi SmsFlashPromotionProductRelationApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      分页查询不同场次关联及商品信息
// @Description  分页查询不同场次关联及商品信息
// @Tags         限时购和商品关系管理
// @Accept       json
// @Produce      json
// @Param        flashPromotionId   query      string  false  "flashPromotionId"
// @Param        flashPromotionSessionId   query      string  false  "flashPromotionSessionId"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashProductRelation/list [get]
func (iApi SmsFlashPromotionProductRelationApi) List(ctx *gin.Context) {
	flashPromotionId := ctx.Query("flashPromotionId")
	flashPromotionSessionId := ctx.Query("flashPromotionSessionId")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(flashPromotionId, flashPromotionSessionId, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
