// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCmsPrefrenceAreaProductRelation = "cms_prefrence_area_product_relation"

// CmsPrefrenceAreaProductRelation mapped from table <cms_prefrence_area_product_relation>
type CmsPrefrenceAreaProductRelation struct {
	ID              int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	PrefrenceAreaID int64 `gorm:"column:prefrence_area_id" json:"prefrenceAreaId" form:"prefrenceAreaId"`
	ProductID       int64 `gorm:"column:product_id" json:"productId" form:"productId"`
}

// TableName CmsPrefrenceAreaProductRelation's table name
func (*CmsPrefrenceAreaProductRelation) TableName() string {
	return TableNameCmsPrefrenceAreaProductRelation
}
