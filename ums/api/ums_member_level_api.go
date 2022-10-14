package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 会员等级管理
type UmsMemberLevelApi struct {
	Service service.UmsMemberLevelService
}

// @Summary      列表
// @Description  列表
// @Tags         会员等级管理
// @Accept       json
// @Produce      json
// @Param        defaultStatus   query      string  false  "defaultStatus"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /memberLevel/list [get]
func (iApi UmsMemberLevelApi) List(ctx *gin.Context) {
	defaultStatus := ctx.Query("defaultStatus")
	list := iApi.Service.List(defaultStatus)
	ctx.JSON(http.StatusOK, util.Data(list))
}
