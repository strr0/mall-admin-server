package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 首页轮播广告管理
type SmsHomeAdvertiseApi struct {
	Service service.SmsHomeAdvertiseService
}

// @Summary      添加广告
// @Description  添加广告
// @Tags         首页轮播广告管理
// @Accept       json
// @Produce      json
// @Param        smsHomeAdvertise   query      model.SmsHomeAdvertise  false  "广告"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/advertise/create [post]
func (iApi SmsHomeAdvertiseApi) Create(ctx *gin.Context) {
	var smsHomeAdvertise model.SmsHomeAdvertise
	_ = ctx.Bind(&smsHomeAdvertise)
	err := iApi.Service.Create(smsHomeAdvertise)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      删除广告
// @Description  删除广告
// @Tags         首页轮播广告管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/advertise/delete/{id} [post]
func (iApi SmsHomeAdvertiseApi) Delete(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      修改上下线状态
// @Description  修改上下线状态
// @Tags         首页轮播广告管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        status   query      string  false  "status"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/advertise/update/status/{id} [post]
func (iApi SmsHomeAdvertiseApi) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.PostForm("status")
	err := iApi.Service.UpdateStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      获取广告详情
// @Description  获取广告详情
// @Tags         首页轮播广告管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/advertise/{id} [get]
func (iApi SmsHomeAdvertiseApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      修改广告
// @Description  修改广告
// @Tags         首页轮播广告管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        smsHomeAdvertise   query      model.SmsHomeAdvertise  false  "广告"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/advertise/update/{id} [post]
func (iApi SmsHomeAdvertiseApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var smsHomeAdvertise model.SmsHomeAdvertise
	_ = ctx.Bind(&smsHomeAdvertise)
	err := iApi.Service.Update(id, smsHomeAdvertise)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      分页查询广告
// @Description  分页查询广告
// @Tags         首页轮播广告管理
// @Accept       json
// @Produce      json
// @Param        name   query      string  false  "名称"
// @Param        type   query      string  false  "类型"
// @Param        endTime   query      string  false  "时间"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/advertise/list [get]
func (iApi SmsHomeAdvertiseApi) List(ctx *gin.Context) {
	name := ctx.Query("name")
	type_ := ctx.Query("type")
	endTime := ctx.Query("endTime")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(name, type_, endTime, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
