package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service/dto"
	"time"
)

// 限购场次
type SmsFlashPromotionSessionService struct {
	DB *gorm.DB
}

func (iService SmsFlashPromotionSessionService) Create(smsFlashPromotionSession model.SmsFlashPromotionSession) error {
	smsFlashPromotionSession.CreateTime = time.Now()
	result := iService.DB.Create(&smsFlashPromotionSession)
	return result.Error
}

func (iService SmsFlashPromotionSessionService) Update(idStr string, smsFlashPromotionSession model.SmsFlashPromotionSession) error {
	result := iService.DB.Save(&smsFlashPromotionSession)
	return result.Error
}

func (iService SmsFlashPromotionSessionService) UpdateStatus(id, status string) error {
	result := iService.DB.Model(&model.SmsFlashPromotionSession{}).Where("id = ?", id).Update("status", status)
	return result.Error
}

func (iService SmsFlashPromotionSessionService) Delete(id string) error {
	result := iService.DB.Delete(&model.SmsFlashPromotionSession{}, id)
	return result.Error
}

func (iService SmsFlashPromotionSessionService) GetItem(id string) *model.SmsFlashPromotionSession {
	var smsFlashPromotionSession model.SmsFlashPromotionSession
	result := iService.DB.First(&smsFlashPromotionSession, id)
	if result.Error != nil {
		return nil
	}
	return &smsFlashPromotionSession
}

func (iService SmsFlashPromotionSessionService) ListAll() []model.SmsFlashPromotionSession {
	var list []model.SmsFlashPromotionSession
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

func (iService SmsFlashPromotionSessionService) SelectList(flashPromotionId string) []dto.SmsFlashPromotionSessionDto {
	dtos := make([]dto.SmsFlashPromotionSessionDto, 0)
	var list []model.SmsFlashPromotionSession
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	for _, flashPromotionSession := range list {
		var count int64
		result := iService.DB.Model(&model.SmsFlashPromotionProductRelation{}).Where("flash_promotion_id = ? and flash_promotion_session_id = ?", flashPromotionId, flashPromotionSession.ID).Count(&count)
		if result.Error == nil {
			dtos = append(dtos, dto.SmsFlashPromotionSessionDto{
				SmsFlashPromotionSession: flashPromotionSession,
				ProductCount:             count,
			})
		}
	}
	return dtos
}
