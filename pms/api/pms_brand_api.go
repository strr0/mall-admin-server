package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service"
	"mall-admin-server/util"
	"net/http"
)

// 商品品牌管理
type PmsBrandApi struct {
	Service service.PmsBrandService
}

// @Summary      获取全部品牌列表
// @Description  获取全部品牌列表
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/listAll [get]
func (iApi PmsBrandApi) ListAll(ctx *gin.Context) {
	all := iApi.Service.ListAll()
	ctx.JSON(http.StatusOK, util.Data(all))
}

// @Summary      添加品牌
// @Description  添加品牌
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        pmsBrand   query      model.PmsBrand  false  "品牌"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/create [post]
func (iApi PmsBrandApi) Create(ctx *gin.Context) {
	var pmsBrand model.PmsBrand
	_ = ctx.Bind(&pmsBrand)
	err := iApi.Service.Create(pmsBrand)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("保存失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("保存成功"))
}

// @Summary      更新品牌
// @Description  更新品牌
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Param        pmsBrand   query      model.PmsBrand  false  "品牌"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/update/{id} [post]
func (iApi PmsBrandApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var pmsBrand model.PmsBrand
	_ = ctx.Bind(&pmsBrand)
	err := iApi.Service.Update(id, pmsBrand)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      删除品牌
// @Description  删除品牌
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/delete/{id} [post]
func (iApi PmsBrandApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := iApi.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      根据品牌名称分页获取品牌列表
// @Description  根据品牌名称分页获取品牌列表
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        keyword   query      string  false  "名称"
// @Param        pageNum   query      string  false  "页码"
// @Param        pageSize   query      string  false  "数量"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/list [get]
func (iApi PmsBrandApi) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	list, total := iApi.Service.List(keyword, pageNum, pageSize)
	ctx.JSON(http.StatusOK, util.Page(list, total))
}

// @Summary      根据编号查询品牌信息
// @Description  根据编号查询品牌信息
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        id   path      string  false  "id"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/{id} [get]
func (iApi PmsBrandApi) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	item := iApi.Service.GetItem(id)
	ctx.JSON(http.StatusOK, util.Data(item))
}

// @Summary      批量删除品牌
// @Description  批量删除品牌
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/delete/batch [post]
func (iApi PmsBrandApi) DeleteBrand(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	err := iApi.Service.DeleteBrand(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Failed("删除失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("删除成功"))
}

// @Summary      批量更新显示状态
// @Description  批量更新显示状态
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        showStatus   query      string  false  "showStatus"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/update/showStatus [post]
func (iApi PmsBrandApi) UpdateShowStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	showStatus := ctx.PostForm("showStatus")
	err := iApi.Service.UpdateShowStatus(ids, showStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Success("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}

// @Summary      批量更新厂家制造商状态
// @Description  批量更新厂家制造商状态
// @Tags         商品品牌管理
// @Accept       json
// @Produce      json
// @Param        ids   query      []string  false  "ids"
// @Param        factoryStatus   query      string  false  "factoryStatus"
// @Success      200  {object}  util.CommonResult
// @Failure      400  {object}  util.CommonResult
// @Failure      404  {object}  util.CommonResult
// @Failure      500  {object}  util.CommonResult
// @Router       /brand/update/factoryStatus [post]
func (iApi PmsBrandApi) UpdateFactoryStatus(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	factoryStatus := ctx.PostForm("factoryStatus")
	err := iApi.Service.UpdateFactoryStatus(ids, factoryStatus)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Success("修改失败"))
		return
	}
	ctx.JSON(http.StatusOK, util.Success("修改成功"))
}
