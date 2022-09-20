package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

type PmsSkuStockApi struct {
	Service service.PmsSkuStockService
}

func (iApi PmsSkuStockApi) GetList(ctx *gin.Context) {
	pid := ctx.Param("pid")
	keyword := ctx.Query("keyword")
	list := iApi.Service.GetList(pid, keyword)
	ctx.JSON(http.StatusOK, util.Data(list))
}

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
