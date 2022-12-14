// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsCouponProductRelation = "sms_coupon_product_relation"

// SmsCouponProductRelation mapped from table <sms_coupon_product_relation>
type SmsCouponProductRelation struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	CouponID    int64  `gorm:"column:coupon_id" json:"couponId" form:"couponId"`
	ProductID   int64  `gorm:"column:product_id" json:"productId" form:"productId"`
	ProductName string `gorm:"column:product_name" json:"productName" form:"productName"` // 商品名称
	ProductSn   string `gorm:"column:product_sn" json:"productSn" form:"productSn"`       // 商品编码
}

// TableName SmsCouponProductRelation's table name
func (*SmsCouponProductRelation) TableName() string {
	return TableNameSmsCouponProductRelation
}
