package dto

import "mall-admin-server/pms/model"

type PmsProductCategoryDto struct {
	model.PmsProductCategory
	ProductAttributeIdList []int64 `form:"productAttributeIdList"`
}
