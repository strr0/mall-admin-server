package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service"
	"mall-admin-server/util"
	"net/http"
)

// 后台资源分类管理
type UmsResourceCategoryApi struct {
	Service service.UmsResourceCategoryService
}

// @Summary      资源分类列表
// @Description  资源分类列表
// @Tags         后台资源分类管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resourceCategory/listAll [get]
func (iApi UmsResourceCategoryApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// @Summary      创建资源分类
// @Description  创建资源分类
// @Tags         后台资源分类管理
// @Accept       json
// @Produce      json
// @Param        umsResourceCategory   query      model.UmsResourceCategory  false  "分类"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resourceCategory/create [post]
func (iApi UmsResourceCategoryApi) Create(ctx *gin.Context) {
	var umsResourceCategory model.UmsResourceCategory
	_ = ctx.Bind(&umsResourceCategory)
	err := iApi.Service.Create(umsResourceCategory)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      修改资源分类
// @Description  修改资源分类
// @Tags         后台资源分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        umsResourceCategory   query      model.UmsResourceCategory  false  "分类"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resourceCategory/update/{id} [post]
func (iApi UmsResourceCategoryApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var umsResourceCategory model.UmsResourceCategory
	_ = ctx.Bind(&umsResourceCategory)
	err := iApi.Service.Update(id, umsResourceCategory)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      修改资源分类
// @Description  修改资源分类
// @Tags         后台资源分类管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /resourceCategory/delete/{id} [post]
func (iApi UmsResourceCategoryApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}
