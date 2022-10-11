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

// @Summary      添加活动
// @Description  添加活动
// @Tags         限时购活动管理
// @Accept       json
// @Produce      json
// @Param        smsFlashPromotion   query      model.SmsFlashPromotion  false  "活动"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flash/create [post]
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

// @Summary      编辑活动
// @Description  编辑活动
// @Tags         限时购活动管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        smsFlashPromotion   query      model.SmsFlashPromotion  false  "活动"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flash/update/{id} [post]
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

// @Summary      删除活动
// @Description  删除活动
// @Tags         限时购活动管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flash/delete/{id} [post]
func (iApi SmsFlashPromotionApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      修改上下线状态
// @Description  修改上下线状态
// @Tags         限时购活动管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        status   query      string  false  "status"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flash/update/status/{id} [post]
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

// @Summary      获取活动详情
// @Description  获取活动详情
// @Tags         限时购活动管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flash/{id} [get]
func (iApi SmsFlashPromotionApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      根据活动名称分页查询
// @Description  根据活动名称分页查询
// @Tags         限时购活动管理
// @Accept       json
// @Produce      json
// @Param        keyword   query      string  false  "名称"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flash/list [get]
func (iApi SmsFlashPromotionApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
