package dto

type PmsProductAttributeDto struct {
	AttributeId         int64 `gorm:"column:attribute_id" json:"attributeId" form:"attributeId"`
	AttributeCategoryId int64 `gorm:"column:attribute_category_id" json:"attributeCategoryId" form:"attributeCategoryId"`
}
