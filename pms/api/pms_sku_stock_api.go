package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

// 商品SKU库存管理
type PmsSkuStockApi struct {
	Service service.PmsSkuStockService
}

// @Summary      根据商品ID及sku编码模糊搜索sku库存
// @Description  根据商品ID及sku编码模糊搜索sku库存
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        pid   path      string  false  "pid"
// @Param        keyword   query      string  false  "keyword"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /sku/{pid} [get]
func (iApi PmsSkuStockApi) GetList(ctx *gin.Context) {
	pid := ctx.Param("pid")
	keyword := ctx.Query("keyword")
	list := iApi.Service.GetList(pid, keyword)
	ctx.JSON(http.StatusOK, util.Data(list))
}

// @Summary      批量更新sku库存信息
// @Description  批量更新sku库存信息
// @Tags         商品分类管理
// @Accept       json
// @Produce      json
// @Param        skuStockList   body      []model.PmsSkuStock  false  "skuStockList"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /sku/update [post]
func (iApi PmsSkuStockApi) Update(ctx *gin.Context) {
	skuStockList := make([]model.PmsSkuStock, 0)
	_ = ctx.Bind(&skuStockList)
	err := iApi.Service.Update(skuStockList)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
