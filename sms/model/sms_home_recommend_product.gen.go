// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsHomeRecommendProduct = "sms_home_recommend_product"

// SmsHomeRecommendProduct mapped from table <sms_home_recommend_product>
type SmsHomeRecommendProduct struct {
	ID              int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	ProductID       int64  `gorm:"column:product_id" json:"productId" form:"productId"`
	ProductName     string `gorm:"column:product_name" json:"productName" form:"productName"`
	RecommendStatus int32  `gorm:"column:recommend_status" json:"recommendStatus" form:"recommendStatus"`
	Sort            int32  `gorm:"column:sort" json:"sort" form:"sort"`
}

// TableName SmsHomeRecommendProduct's table name
func (*SmsHomeRecommendProduct) TableName() string {
	return TableNameSmsHomeRecommendProduct
}
