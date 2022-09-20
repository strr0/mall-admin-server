package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 限时购场次管理
type SmsFlashPromotionSessionApi struct {
	Service service.SmsFlashPromotionSessionService
}

// 添加场次
func (iApi SmsFlashPromotionSessionApi) Create(ctx *gin.Context) {
	var smsFlashPromotionSession model.SmsFlashPromotionSession
	_ = ctx.Bind(&smsFlashPromotionSession)
	err := iApi.Service.Create(smsFlashPromotionSession)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 修改场次
func (iApi SmsFlashPromotionSessionApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var smsFlashPromotionSession model.SmsFlashPromotionSession
	_ = ctx.Bind(&smsFlashPromotionSession)
	err := iApi.Service.Update(id, smsFlashPromotionSession)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 修改启用状态
func (iApi SmsFlashPromotionSessionApi) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.PostForm("status")
	err := iApi.Service.UpdateStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 删除场次
func (iApi SmsFlashPromotionSessionApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 获取场次详情
func (iApi SmsFlashPromotionSessionApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 获取全部场次
func (iApi SmsFlashPromotionSessionApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// 获取全部可选场次及其数量
func (iApi SmsFlashPromotionSessionApi) SelectList(ctx *gin.Context) {
	flashPromotionId := ctx.Query("flashPromotionId")
	list := iApi.Service.SelectList(flashPromotionId)
	ctx.JSON(http.StatusOK, util.Data(list))
}
