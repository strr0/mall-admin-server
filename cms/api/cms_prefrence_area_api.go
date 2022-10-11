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

// @Summary      获取所有商品优选
// @Description  获取所有商品优选
// @Tags         商品优选管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /prefrenceArea/listAll [get]
func (iApi CmsPrefrenceAreaApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}
