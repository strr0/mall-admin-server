package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 限时购活动管理
type SmsFlashPromotionApi struct {
	Service service.SmsFlashPromotionService
}

// 添加活动
func (iApi SmsFlashPromotionApi) Create(ctx *gin.Context) {
	var smsFlashPromotion model.SmsFlashPromotion
	_ = ctx.Bind(&smsFlashPromotion)
	err := iApi.Service.Create(smsFlashPromotion)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 编辑活动
func (iApi SmsFlashPromotionApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var smsFlashPromotion model.SmsFlashPromotion
	_ = ctx.Bind(&smsFlashPromotion)
	err := iApi.Service.Update(id, smsFlashPromotion)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 删除活动
func (iApi SmsFlashPromotionApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 修改上下线状态
func (iApi SmsFlashPromotionApi) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.PostForm("status")
	err := iApi.Service.UpdateStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 获取活动详情
func (iApi SmsFlashPromotionApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 根据活动名称分页查询
func (iApi SmsFlashPromotionApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
