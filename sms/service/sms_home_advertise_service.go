package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/util"
)

// 首页广告
type SmsHomeAdvertiseService struct {
	DB *gorm.DB
}

func (iService SmsHomeAdvertiseService) Create(smsHomeAdvertise model.SmsHomeAdvertise) error {
	smsHomeAdvertise.ClickCount = 0
	smsHomeAdvertise.OrderCount = 0
	result := iService.DB.Create(&smsHomeAdvertise)
	return result.Error
}

func (iService SmsHomeAdvertiseService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.SmsHomeAdvertise{}, ids)
	return result.Error
}

func (iService SmsHomeAdvertiseService) UpdateStatus(id, status string) error {
	result := iService.DB.Model(&model.SmsHomeAdvertise{}).Where("id = ?", id).Update("status", status)
	return result.Error
}

func (iService SmsHomeAdvertiseService) GetItem(id string) *model.SmsHomeAdvertise {
	var smsHomeAdvertise model.SmsHomeAdvertise
	result := iService.DB.First(&smsHomeAdvertise, id)
	if result.Error != nil {
		return nil
	}
	return &smsHomeAdvertise
}

func (iService SmsHomeAdvertiseService) Update(idStr string, smsHomeAdvertise model.SmsHomeAdvertise) error {
	result := iService.DB.Save(&smsHomeAdvertise)
	return result.Error
}

func (iService SmsHomeAdvertiseService) List(name, type_, endTime, pageStr, sizeStr string) ([]model.SmsHomeAdvertise, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if name != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("name like ?", "%"+name+"%")
		})
	}
	if type_ != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("type = ?", type_)
		})
	}
	if endTime != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("end_time >= ? and end_time <= ?", endTime+" 00:00:00", endTime+" 23:59:59")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.SmsHomeAdvertise{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.SmsHomeAdvertise
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
