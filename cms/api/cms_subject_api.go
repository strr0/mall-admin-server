package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/cms/service"
	"mall-admin-server/util"
	"net/http"
)

// 商品专题管理
type CmsSubjectApi struct {
	Service service.CmsSubjectService
}

func (iApi CmsSubjectApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

func (iApi CmsSubjectApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}
