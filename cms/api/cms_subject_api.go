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

// @Summary      获取全部商品专题
// @Description  获取全部商品专题
// @Tags         商品专题管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /subject/listAll [get]
func (iApi CmsSubjectApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// @Summary      根据专题名称分页获取商品专题
// @Description  根据专题名称分页获取商品专题
// @Tags         商品专题管理
// @Accept       json
// @Produce      json
// @Param        keyword   query      string  false  "keyword"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /subject/list [get]
func (iApi CmsSubjectApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, count := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, count))
}
