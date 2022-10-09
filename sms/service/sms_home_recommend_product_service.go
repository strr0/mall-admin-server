package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/util"
)

// 首页人气推荐管理
type SmsHomeRecommendProductService struct {
	DB *gorm.DB
}

func (iService SmsHomeRecommendProductService) Create(smsHomeRecommendProduct model.SmsHomeRecommendProduct) error {
	smsHomeRecommendProduct.RecommendStatus = 1
	smsHomeRecommendProduct.Sort = 0
	result := iService.DB.Create(&smsHomeRecommendProduct)
	return result.Error
}

func (iService SmsHomeRecommendProductService) UpdateSort(id, sort string) error {
	result := iService.DB.Model(&model.SmsHomeRecommendProduct{}).Where("id = ?", id).Update("sort", sort)
	return result.Error
}

func (iService SmsHomeRecommendProductService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.SmsHomeRecommendProduct{}, ids)
	return result.Error
}

func (iService SmsHomeRecommendProductService) UpdateRecommendStatus(ids []string, recommendStatus string) error {
	result := iService.DB.Model(&model.SmsHomeRecommendProduct{}).Where("id in ?", ids).Update("recommend_status", recommendStatus)
	return result.Error
}

func (iService SmsHomeRecommendProductService) List(productName, recommendStatus, pageStr, sizeStr string) ([]model.SmsHomeRecommendProduct, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if productName != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("product_name like ?", "%"+productName+"%")
		})
	}
	if recommendStatus != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("recommend_status = ?", recommendStatus)
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.SmsHomeRecommendProduct{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.SmsHomeRecommendProduct
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
