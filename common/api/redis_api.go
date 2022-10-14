package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/common/service"
	"mall-admin-server/util"
	"net/http"
)

type RedisApi struct {
	Service service.RedisService
}

// @Summary      redis get
// @Description  redis get
// @Tags         Redis
// @Accept       json
// @Produce      json
// @Param        key   query      string  false  "key"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /redis/get [get]
func (iApi RedisApi) Get(ctx *gin.Context) {
	key := ctx.Query("key")
	value := iApi.Service.Get(key)
	if value != "" {
		ctx.JSON(http.StatusOK, util.Data(value))
		return
	}
	ctx.JSON(http.StatusOK, util.Failed("获取失败"))
}

// @Summary      redis set
// @Description  redis set
// @Tags         Redis
// @Accept       json
// @Produce      json
// @Param        key   query      string  false  "key"
// @Param        value   query      string  false  "value"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /redis/set [get]
func (iApi RedisApi) Set(ctx *gin.Context) {
	key := ctx.Query("key")
	value := ctx.Query("value")
	err := iApi.Service.Set(key, value)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("操作失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("操作成功"))
}
