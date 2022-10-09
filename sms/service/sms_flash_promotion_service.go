package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/util"
)

// 限购活动
type SmsFlashPromotionService struct {
	DB *gorm.DB
}

func (iService SmsFlashPromotionService) Create(smsFlashPromotion model.SmsFlashPromotion) error {
	result := iService.DB.Create(&smsFlashPromotion)
	return result.Error
}

func (iService SmsFlashPromotionService) Update(idStr string, smsFlashPromotion model.SmsFlashPromotion) error {
	result := iService.DB.Save(&smsFlashPromotion)
	return result.Error
}

func (iService SmsFlashPromotionService) Delete(id string) error {
	result := iService.DB.Delete(&model.SmsFlashPromotion{}, id)
	return result.Error
}

func (iService SmsFlashPromotionService) UpdateStatus(id, status string) error {
	result := iService.DB.Model(&model.SmsFlashPromotion{}).Where("id = ?", id).Update("status", status)
	return result.Error
}

func (iService SmsFlashPromotionService) GetItem(id string) *model.SmsFlashPromotion {
	var smsFlashPromotion model.SmsFlashPromotion
	result := iService.DB.First(&smsFlashPromotion, id)
	if result.Error != nil {
		return nil
	}
	return &smsFlashPromotion
}

func (iService SmsFlashPromotionService) List(keyword, pageStr, sizeStr string) ([]model.SmsFlashPromotion, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if keyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("title like ?", "%"+keyword+"%")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.SmsFlashPromotion{}).Count(&count)
	if result.Error != nil || count == 0 {
		return nil, 0
	}
	var list []model.SmsFlashPromotion
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
