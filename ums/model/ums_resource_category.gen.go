// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUmsResourceCategory = "ums_resource_category"

// UmsResourceCategory mapped from table <ums_resource_category>
type UmsResourceCategory struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime" form:"createTime"` // 创建时间
	Name       string    `gorm:"column:name" json:"name" form:"name"`                    // 分类名称
	Sort       int32     `gorm:"column:sort" json:"sort" form:"sort"`                    // 排序
}

// TableName UmsResourceCategory's table name
func (*UmsResourceCategory) TableName() string {
	return TableNameUmsResourceCategory
}