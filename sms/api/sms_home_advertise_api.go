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

// 添加广告
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

// 删除广告
func (iApi SmsHomeAdvertiseApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 修改上下线状态
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

// 获取广告详情
func (iApi SmsHomeAdvertiseApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// 修改广告
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

// 分页查询广告
func (iApi SmsHomeAdvertiseApi) List(ctx *gin.Context) {
	name := ctx.Query("name")
	type_ := ctx.Query("type")
	endTime := ctx.Query("endTime")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(name, type_, endTime, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
