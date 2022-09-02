package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

type UmsMemberLevelApi struct {
	Service service.UmsMemberLevelService
}

// 会员等级管理
func (iApi UmsMemberLevelApi) List(ctx *gin.Context) {
	defaultStatus := ctx.Query("defaultStatus")
	list := iApi.Service.List(defaultStatus)
	ctx.JSON(http.StatusOK, util.Data(list))
}
