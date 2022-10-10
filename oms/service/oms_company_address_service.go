package service

import (
	"gorm.io/gorm"
	"mall-admin-server/oms/model"
)

// 收货地址管理
type OmsCompanyAddressService struct {
	DB *gorm.DB
}

func (iService OmsCompanyAddressService) List() []model.OmsCompanyAddress {
	var list []model.OmsCompanyAddress
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}
