package dto

import "mall-admin-server/pms/model"

type PmsProductAttributeCategoryDto struct {
	model.PmsProductAttributeCategory
	ProductAttributeList []model.PmsProductAttribute `json:"productAttributeList"`
}
