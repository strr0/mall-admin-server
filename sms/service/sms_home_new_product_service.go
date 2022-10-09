package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/util"
)

// 首页新品推荐管理
type SmsHomeNewProductService struct {
	DB *gorm.DB
}

func (iService SmsHomeNewProductService) Create(smsHomeNewProduct model.SmsHomeNewProduct) error {
	smsHomeNewProduct.RecommendStatus = 1
	smsHomeNewProduct.Sort = 0
	result := iService.DB.Create(&smsHomeNewProduct)
	return result.Error
}

func (iService SmsHomeNewProductService) UpdateSort(id, sort string) error {
	result := iService.DB.Model(&model.SmsHomeNewProduct{}).Where("id = ?", id).Update("sort", sort)
	return result.Error
}

func (iService SmsHomeNewProductService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.SmsHomeNewProduct{}, ids)
	return result.Error
}

func (iService SmsHomeNewProductService) UpdateRecommendStatus(ids []string, recommendStatus string) error {
	result := iService.DB.Model(&model.SmsHomeNewProduct{}).Where("id in ?", ids).Update("recommend_status", recommendStatus)
	return result.Error
}

func (iService SmsHomeNewProductService) List(productName, recommendStatus, pageStr, sizeStr string) ([]model.SmsHomeNewProduct, int64) {
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
	result := scopes.Model(&model.SmsHomeNewProduct{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.SmsHomeNewProduct
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
