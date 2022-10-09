package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/util"
)

// 首页品牌管理
type SmsHomeBrandService struct {
	DB *gorm.DB
}

func (iService SmsHomeBrandService) Create(smsHomeBrand model.SmsHomeBrand) error {
	smsHomeBrand.RecommendStatus = 1
	smsHomeBrand.Sort = 0
	result := iService.DB.Create(&smsHomeBrand)
	return result.Error
}

func (iService SmsHomeBrandService) UpdateSort(id, sort string) error {
	result := iService.DB.Model(&model.SmsHomeBrand{}).Where("id = ?", id).Update("sort", sort)
	return result.Error
}

func (iService SmsHomeBrandService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.SmsHomeBrand{}, ids)
	return result.Error
}

func (iService SmsHomeBrandService) UpdateRecommendStatus(ids []string, recommendStatus string) error {
	result := iService.DB.Model(&model.SmsHomeBrand{}).Where("id in ?", ids).Update("recommend_status", recommendStatus)
	return result.Error
}

func (iService SmsHomeBrandService) List(brandName, recommendStatus, pageStr, sizeStr string) ([]model.SmsHomeBrand, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if brandName != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("brand_name like ?", "%"+brandName+"%")
		})
	}
	if recommendStatus != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("recommend_status = ?", recommendStatus)
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.SmsHomeBrand{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.SmsHomeBrand
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
