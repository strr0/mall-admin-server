package service

import (
	"gorm.io/gorm"
	"mall-admin-server/oms/model"
)

// 订单设置管理
type OmsOrderSettingService struct {
	DB *gorm.DB
}

func (iService OmsOrderSettingService) GetItem(id string) *model.OmsOrderSetting {
	var omsOrderSetting model.OmsOrderSetting
	result := iService.DB.First(&omsOrderSetting, id)
	if result.Error != nil {
		return nil
	}
	return &omsOrderSetting
}

func (iService OmsOrderSettingService) Update(idStr string, omsOrderSetting model.OmsOrderSetting) error {
	result := iService.DB.Save(&omsOrderSetting)
	return result.Error
}
