package service

import (
	"gorm.io/gen"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/query"
	"mall-admin-server/util"
)

// 商品SKU库存管理
type PmsSkuStockService struct {
	//
}

func (PmsSkuStockService) GetList(pidStr, keyword string) []*model.PmsSkuStock {
	pid, err := util.ParseInt64WithErr(pidStr)
	if err != nil {
		return nil
	}
	pmsSkuStock := query.PmsSkuStock
	conds := make([]gen.Condition, 0)
	conds = append(conds, query.PmsSkuStock.ProductID.Eq(pid))
	if keyword != "" {
		conds = append(conds, query.PmsSkuStock.SkuCode.Like("%"+keyword+"%"))
	}
	find, err := pmsSkuStock.Where(conds...).Find()
	if err != nil {
		return nil
	}
	return find
}

func (PmsSkuStockService) Update(pmsSkuStocks []model.PmsSkuStock) error {
	for _, pmsSkuStock := range pmsSkuStocks {
		_ = query.PmsSkuStock.Create(&pmsSkuStock)
	}
	return nil
}
