package dto

import "mall-admin-server/pms/model"

type PmsProductCategoryTreeDto struct {
	model.PmsProductCategory
	Children []PmsProductCategoryTreeDto `json:"children"`
}
