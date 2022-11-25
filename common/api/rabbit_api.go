package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/common/service"
	"mall-admin-server/util"
	"net/http"
)

type RabbitApi struct {
	Service service.RabbitService
}

// @Summary      rabbit 发送消息
// @Description  rabbit 发送消息
// @Tags         rabbit
// @Accept       json
// @Produce      json
// @Param        msg   query      string  false  "消息"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /rabbit/send [post]
func (iApi RabbitApi) Send(ctx *gin.Context) {
	msg := ctx.Query("msg")
	err := iApi.Service.Send(msg)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("发送失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("发送成功"))
}

// @Summary      rabbit 接收消息
// @Description  rabbit 接收消息
// @Tags         rabbit
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /rabbit/recv [get]
func (iApi RabbitApi) Recv(ctx *gin.Context) {
	go iApi.Service.Recv()
	ctx.JSON(http.StatusOK, util.Success("操作成功"))
}

// @Summary      rabbit 初始化队列
// @Description  rabbit 初始化队列
// @Tags         rabbit
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /rabbit/initQueue [get]
func (iApi RabbitApi) InitQueue(ctx *gin.Context) {
	iApi.Service.InitQueue()
	ctx.JSON(http.StatusOK, util.Success("操作成功"))
}
