package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service/dto"
	"mall-admin-server/util"
)

type SmsFlashPromotionProductRelationService struct {
	DB *gorm.DB
}

func (iService SmsFlashPromotionProductRelationService) Create(smsFlashPromotionProductRelation model.SmsFlashPromotionProductRelation) error {
	result := iService.DB.Create(&smsFlashPromotionProductRelation)
	return result.Error
}

func (iService SmsFlashPromotionProductRelationService) Update(idStr string, smsFlashPromotionProductRelation model.SmsFlashPromotionProductRelation) error {
	result := iService.DB.Save(&smsFlashPromotionProductRelation)
	return result.Error
}

func (iService SmsFlashPromotionProductRelationService) Delete(id string) error {
	result := iService.DB.Delete(&model.SmsFlashPromotionProductRelation{}, id)
	return result.Error
}

func (iService SmsFlashPromotionProductRelationService) GetItem(id string) *model.SmsFlashPromotionProductRelation {
	var smsFlashPromotionProductRelation model.SmsFlashPromotionProductRelation
	result := iService.DB.First(&smsFlashPromotionProductRelation, id)
	if result.Error != nil {
		return nil
	}
	return &smsFlashPromotionProductRelation
}

func (iService SmsFlashPromotionProductRelationService) List(flashPromotionId, flashPromotionSessionId, pageStr, sizeStr string) ([]dto.SmsFlashPromotionProductRelationDto, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	raw := iService.DB.Raw("SELECT r.id, r.flash_promotion_price, r.flash_promotion_count, r.flash_promotion_limit, r.flash_promotion_id, r.flash_promotion_session_id, r.product_id, r.sort, p.`name` name, p.product_sn product_sn, p.price price, p.stock stock FROM sms_flash_promotion_product_relation r LEFT JOIN pms_product p ON r.product_id = p.id WHERE r.flash_promotion_id = ? AND r.flash_promotion_session_id = ? ORDER BY r.sort DESC", flashPromotionId, flashPromotionSessionId)
	var count int64
	result := raw.Count(&count)
	if result.Error != nil || count == 0 {
		return nil, 0
	}
	var list []dto.SmsFlashPromotionProductRelationDto
	result = raw.Offset(offset).Limit(size).Scan(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService SmsFlashPromotionProductRelationService) GetCount(flashPromotionId, flashPromotionSessionId string) int64 {
	var count int64
	result := iService.DB.Where("flash_promotion_id = ? and flash_promotion_session_id = ?", flashPromotionId, flashPromotionSessionId).Count(&count)
	if result.Error != nil {
		return 0
	}
	return count
}
