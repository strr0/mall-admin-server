package dto

import (
	pms "mall-admin-server/pms/model"
	sms "mall-admin-server/sms/model"
)

type SmsFlashPromotionProductRelationDto struct {
	sms.SmsFlashPromotionProductRelation
	pms.PmsProduct
}
