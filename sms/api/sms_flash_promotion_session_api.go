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

// @Summary      添加场次
// @Description  添加场次
// @Tags         限时购场次管理
// @Accept       json
// @Produce      json
// @Param        smsFlashPromotionSession   query      model.SmsFlashPromotionSession  false  "场次"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashSession/create [post]
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

// @Summary      修改场次
// @Description  修改场次
// @Tags         限时购场次管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        smsFlashPromotionSession   query      model.SmsFlashPromotionSession  false  "场次"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashSession/update/{id} [post]
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

// @Summary      修改启用状态
// @Description  修改启用状态
// @Tags         限时购场次管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        status   query      string  false  "status"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashSession/update/status/{id} [post]
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

// @Summary      删除场次
// @Description  删除场次
// @Tags         限时购场次管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashSession/delete/{id} [post]
func (iApi SmsFlashPromotionSessionApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      获取场次详情
// @Description  获取场次详情
// @Tags         限时购场次管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashSession/{id} [get]
func (iApi SmsFlashPromotionSessionApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      获取全部场次
// @Description  获取全部场次
// @Tags         限时购场次管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashSession/listAll [get]
func (iApi SmsFlashPromotionSessionApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// @Summary      获取全部可选场次及其数量
// @Description  获取全部可选场次及其数量
// @Tags         限时购场次管理
// @Accept       json
// @Produce      json
// @Param        flashPromotionId   query      string  false  "flashPromotionId"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /flashSession/selectList [get]
func (iApi SmsFlashPromotionSessionApi) SelectList(ctx *gin.Context) {
	flashPromotionId := ctx.Query("flashPromotionId")
	list := iApi.Service.SelectList(flashPromotionId)
	ctx.JSON(http.StatusOK, util.Data(list))
}
