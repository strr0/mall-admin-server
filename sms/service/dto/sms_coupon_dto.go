package dto

import "mall-admin-server/sms/model"

type SmsCouponDto struct {
	model.SmsCoupon
	ProductRelationList         []model.SmsCouponProductRelation
	ProductCategoryRelationList []model.SmsCouponProductCategoryRelation
}
