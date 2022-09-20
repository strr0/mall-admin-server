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

// 批量选择商品添加关联
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

// 修改关联信息
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

// 删除关联
func (iApi SmsFlashPromotionProductRelationApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 获取管理商品促销信息
func (iApi SmsFlashPromotionProductRelationApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 分页查询不同场次关联及商品信息
func (iApi SmsFlashPromotionProductRelationApi) List(ctx *gin.Context) {
	flashPromotionId := ctx.Query("flashPromotionId")
	flashPromotionSessionId := ctx.Query("flashPromotionSessionId")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(flashPromotionId, flashPromotionSessionId, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
