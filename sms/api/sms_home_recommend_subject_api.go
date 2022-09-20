package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service"
	"mall-admin-server/util"
	"net/http"
)

// 首页专题推荐管理
type SmsHomeRecommendSubjectApi struct {
	Service service.SmsHomeRecommendSubjectService
}

// 添加首页推荐专题
func (iApi SmsHomeRecommendSubjectApi) Create(ctx *gin.Context) {
	var smsHomeRecommendSubject model.SmsHomeRecommendSubject
	_ = ctx.Bind(&smsHomeRecommendSubject)
	err := iApi.Service.Create(smsHomeRecommendSubject)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// 修改推荐排序
func (iApi SmsHomeRecommendSubjectApi) UpdateSort(ctx *gin.Context) {
	id := ctx.Param("id")
	sort := ctx.PostForm("sort")
	err := iApi.Service.UpdateSort(id, sort)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 批量删除推荐
func (iApi SmsHomeRecommendSubjectApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// 批量修改推荐状态
func (iApi SmsHomeRecommendSubjectApi) UpdateRecommendStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	recommendStatus := ctx.PostForm("recommendStatus")
	err := iApi.Service.UpdateRecommendStatus(ids, recommendStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// 分页查询推荐
func (iApi SmsHomeRecommendSubjectApi) List(ctx *gin.Context) {
	subjectName := ctx.Query("subjectName")
	recommendStatus := ctx.Query("recommendStatus")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(subjectName, recommendStatus, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
