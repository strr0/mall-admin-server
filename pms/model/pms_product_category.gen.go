// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsProductCategory = "pms_product_category"

// PmsProductCategory mapped from table <pms_product_category>
type PmsProductCategory struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	ParentID     int64  `gorm:"column:parent_id" json:"parentId" form:"parentId"` // 上机分类的编号：0表示一级分类
	Name         string `gorm:"column:name" json:"name" form:"name"`
	Level        int32  `gorm:"column:level" json:"level" form:"level"` // 分类级别：0->1级；1->2级
	ProductCount int32  `gorm:"column:product_count" json:"productCount" form:"productCount"`
	ProductUnit  string `gorm:"column:product_unit" json:"productUnit" form:"productUnit"`
	NavStatus    int32  `gorm:"column:nav_status" json:"navStatus" form:"navStatus"`    // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int32  `gorm:"column:show_status" json:"showStatus" form:"showStatus"` // 显示状态：0->不显示；1->显示
	Sort         int32  `gorm:"column:sort" json:"sort" form:"sort"`
	Icon         string `gorm:"column:icon" json:"icon" form:"icon"` // 图标
	Keywords     string `gorm:"column:keywords" json:"keywords" form:"keywords"`
	Description  string `gorm:"column:description" json:"description" form:"description"` // 描述
}

// TableName PmsProductCategory's table name
func (*PmsProductCategory) TableName() string {
	return TableNamePmsProductCategory
}
