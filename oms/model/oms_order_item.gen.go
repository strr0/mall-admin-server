// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOmsOrderItem = "oms_order_item"

// OmsOrderItem mapped from table <oms_order_item>
type OmsOrderItem struct {
	ID                int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	OrderID           int64   `gorm:"column:order_id" json:"orderId" form:"orderId"` // 订单id
	OrderSn           string  `gorm:"column:order_sn" json:"orderSn" form:"orderSn"` // 订单编号
	ProductID         int64   `gorm:"column:product_id" json:"productId" form:"productId"`
	ProductPic        string  `gorm:"column:product_pic" json:"productPic" form:"productPic"`
	ProductName       string  `gorm:"column:product_name" json:"productName" form:"productName"`
	ProductBrand      string  `gorm:"column:product_brand" json:"productBrand" form:"productBrand"`
	ProductSn         string  `gorm:"column:product_sn" json:"productSn" form:"productSn"`
	ProductPrice      float64 `gorm:"column:product_price" json:"productPrice" form:"productPrice"`                 // 销售价格
	ProductQuantity   int32   `gorm:"column:product_quantity" json:"productQuantity" form:"productQuantity"`        // 购买数量
	ProductSkuID      int64   `gorm:"column:product_sku_id" json:"productSkuId" form:"productSkuId"`                // 商品sku编号
	ProductSkuCode    string  `gorm:"column:product_sku_code" json:"productSkuCode" form:"productSkuCode"`          // 商品sku条码
	ProductCategoryID int64   `gorm:"column:product_category_id" json:"productCategoryId" form:"productCategoryId"` // 商品分类id
	PromotionName     string  `gorm:"column:promotion_name" json:"promotionName" form:"promotionName"`              // 商品促销名称
	PromotionAmount   float64 `gorm:"column:promotion_amount" json:"promotionAmount" form:"promotionAmount"`        // 商品促销分解金额
	CouponAmount      float64 `gorm:"column:coupon_amount" json:"couponAmount" form:"couponAmount"`                 // 优惠券优惠分解金额
	IntegrationAmount float64 `gorm:"column:integration_amount" json:"integrationAmount" form:"integrationAmount"`  // 积分优惠分解金额
	RealAmount        float64 `gorm:"column:real_amount" json:"realAmount" form:"realAmount"`                       // 该商品经过优惠后的分解金额
	GiftIntegration   int32   `gorm:"column:gift_integration" json:"giftIntegration" form:"giftIntegration"`
	GiftGrowth        int32   `gorm:"column:gift_growth" json:"giftGrowth" form:"giftGrowth"`
	ProductAttr       string  `gorm:"column:product_attr" json:"productAttr" form:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

// TableName OmsOrderItem's table name
func (*OmsOrderItem) TableName() string {
	return TableNameOmsOrderItem
}
