package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 会员登录注册管理
type UmsMemberApi struct {
	Service service.UmsMemberService
}

// @Summary      获取验证码
// @Description  获取验证码
// @Tags         会员登录注册管理
// @Accept       json
// @Produce      json
// @Param        telephone   query      string  false  "telephone"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /member/generateAuthCode [get]
func (iApi UmsMemberApi) GenerateAuthCode(ctx *gin.Context) {
	telephone := ctx.Query("telephone")
	err := iApi.Service.GenerateAuthCode(telephone)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("操作失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("操作成功"))
}

// @Summary      验证码校验
// @Description  验证码校验
// @Tags         会员登录注册管理
// @Accept       json
// @Produce      json
// @Param        telephone   query      string  false  "telephone"
// @Param        authCode   query      string  false  "authCode"
// @Security ApiKeyAuth
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /member/verifyAuthCode [get]
func (iApi UmsMemberApi) VerifyAuthCode(ctx *gin.Context) {
	telephone := ctx.Query("telephone")
	authCode := ctx.Query("authCode")
	if iApi.Service.VerifyAuthCode(telephone, authCode) {
		ctx.JSON(http.StatusOK, util.Success("验证成功"))
		return
	}
	ctx.JSON(http.StatusOK, util.Failed("验证失败"))
}
