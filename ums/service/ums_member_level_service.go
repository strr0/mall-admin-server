package service

import (
	"gorm.io/gorm"
	"mall-admin-server/ums/model"
)

type UmsMemberLevelService struct {
	DB *gorm.DB
}

func (iService UmsMemberLevelService) List(defaultStatus string) []model.UmsMemberLevel {
	var list []model.UmsMemberLevel
	iService.DB.Where("default_status = ?", defaultStatus).Find(&list)
	return list
}
