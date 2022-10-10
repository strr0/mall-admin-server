package service

import (
	"gorm.io/gorm"
	"mall-admin-server/cms/model"
)

// 商品优选管理
type CmsPrefrenceAreaService struct {
	DB *gorm.DB
}

func (iService CmsPrefrenceAreaService) ListAll() []model.CmsPrefrenceArea {
	var list []model.CmsPrefrenceArea
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}
