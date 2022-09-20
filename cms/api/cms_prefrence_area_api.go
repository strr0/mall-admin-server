package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/cms/service"
	"mall-admin-server/util"
	"net/http"
)

// 商品优选管理
type CmsPrefrenceAreaApi struct {
	Service service.CmsPrefrenceAreaService
}

func (iApi CmsPrefrenceAreaApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}
