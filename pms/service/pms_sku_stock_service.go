package service

import (
	"gorm.io/gorm"
	"mall-admin-server/pms/model"
)

// 商品SKU库存管理
type PmsSkuStockService struct {
	DB *gorm.DB
}

func (iService PmsSkuStockService) GetList(pid, keyword string) []model.PmsSkuStock {
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where("product_id = ?", pid)
	})
	if keyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("sku_code like ?", "%"+keyword+"%")
		})
	}
	var list []model.PmsSkuStock
	result := iService.DB.Scopes(funcs...).Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

func (iService PmsSkuStockService) Update(pmsSkuStocks []model.PmsSkuStock) error {
	result := iService.DB.Save(&pmsSkuStocks)
	return result.Error
}
