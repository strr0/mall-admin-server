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

// @Summary      添加首页推荐专题
// @Description  添加首页推荐专题
// @Tags         首页专题推荐管理
// @Accept       json
// @Produce      json
// @Param        smsHomeRecommendSubject   query      model.SmsHomeRecommendSubject  false  "首页推荐"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/recommendSubject/create [post]
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

// @Summary      修改推荐排序
// @Description  修改推荐排序
// @Tags         首页专题推荐管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        sort   query      string  false  "sort"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/recommendSubject/update/sort/{id} [post]
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

// @Summary      批量删除推荐
// @Description  批量删除推荐
// @Tags         首页专题推荐管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/recommendSubject/delete [post]
func (iApi SmsHomeRecommendSubjectApi) Delete(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      批量修改推荐状态
// @Description  批量修改推荐状态
// @Tags         首页专题推荐管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        recommendStatus   query      string  false  "recommendStatus"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/recommendSubject/update/recommendStatus [post]
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

// @Summary      分页查询推荐
// @Description  分页查询推荐
// @Tags         首页专题推荐管理
// @Accept       json
// @Produce      json
// @Param        subjectName   query      string  false  "subjectName"
// @Param        recommendStatus   query      string  false  "recommendStatus"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /home/recommendSubject/list [get]
func (iApi SmsHomeRecommendSubjectApi) List(ctx *gin.Context) {
	subjectName := ctx.Query("subjectName")
	recommendStatus := ctx.Query("recommendStatus")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(subjectName, recommendStatus, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}
