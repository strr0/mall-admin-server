// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsProductFullReduction = "pms_product_full_reduction"

// PmsProductFullReduction mapped from table <pms_product_full_reduction>
type PmsProductFullReduction struct {
	ID          int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	ProductID   int64   `gorm:"column:product_id" json:"productId" form:"productId"`
	FullPrice   float64 `gorm:"column:full_price" json:"fullPrice" form:"fullPrice"`
	ReducePrice float64 `gorm:"column:reduce_price" json:"reducePrice" form:"reducePrice"`
}

// TableName PmsProductFullReduction's table name
func (*PmsProductFullReduction) TableName() string {
	return TableNamePmsProductFullReduction
}
