package dto

import "mall-admin-server/sms/model"

type SmsFlashPromotionSessionDto struct {
	model.SmsFlashPromotionSession
	ProductCount int64 `json:"productCount"`
}
