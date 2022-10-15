package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/common/model"
	"mall-admin-server/common/service"
	"mall-admin-server/util"
	"net/http"
)

type MemberReadHistoryApi struct {
	Service service.MemberReadHistoryService
}

// @Summary      创建历史记录
// @Description  创建历史记录
// @Tags         历史记录
// @Accept       json
// @Produce      json
// @Param        history   query      model.MemberReadHistory  false  "历史记录"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /memberReadHistory/create [post]
func (iApi MemberReadHistoryApi) Create(ctx *gin.Context) {
	var history model.MemberReadHistory
	_ = ctx.Bind(&history)
	err := iApi.Service.Create(history)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      删除历史记录
// @Description  删除历史记录
// @Tags         历史记录
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /memberReadHistory/delete [post]
func (iApi MemberReadHistoryApi) Delete(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	err := iApi.Service.Delete(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      获取历史记录
// @Description  获取历史记录
// @Tags         历史记录
// @Accept       json
// @Produce      json
// @Param        memberId   query      string  false  "memberId"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /memberReadHistory/list [get]
func (iApi MemberReadHistoryApi) List(ctx *gin.Context) {
	memberId := ctx.Query("memberId")
	list := iApi.Service.List(memberId)
	ctx.JSON(http.StatusOK, util.Data(list))
}
