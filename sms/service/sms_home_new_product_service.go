package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/util"
)

// 首页新品推荐管理
type SmsHomeNewProductService struct {
	//
}

func (SmsHomeNewProductService) Create(smsHomeNewProduct model.SmsHomeNewProduct) error {
	smsHomeNewProduct.RecommendStatus = 1
	smsHomeNewProduct.Sort = 0
	return query.SmsHomeNewProduct.Create(&smsHomeNewProduct)
}

func (SmsHomeNewProductService) UpdateSort(idStr, sortStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	sort, err := util.ParseInt32WithErr(sortStr)
	if err != nil {
		return err
	}
	var smsHomeNewProduct model.SmsHomeNewProduct
	smsHomeNewProduct.Sort = sort
	_, err = query.SmsHomeNewProduct.Where(query.SmsHomeNewProduct.ID.Eq(id)).Updates(smsHomeNewProduct)
	return err
}

func (SmsHomeNewProductService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.SmsHomeNewProduct.Where(query.SmsHomeNewProduct.ID.In(ids...)).Delete()
	return err
}

func (SmsHomeNewProductService) UpdateRecommendStatus(idsStr []string, recommendStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	recommendStatus, err := util.ParseInt32WithErr(recommendStatusStr)
	if err != nil {
		return err
	}
	var smsHomeNewProduct model.SmsHomeNewProduct
	smsHomeNewProduct.RecommendStatus = recommendStatus
	_, err = query.SmsHomeNewProduct.Where(query.SmsHomeNewProduct.ID.In(ids...)).Updates(smsHomeNewProduct)
	return err
}

func (SmsHomeNewProductService) List(productName, recommendStatusStr, pageStr, sizeStr string) ([]*model.SmsHomeNewProduct, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsHomeNewProduct := query.SmsHomeNewProduct
	conds := make([]gen.Condition, 0)
	if productName != "" {
		conds = append(conds, smsHomeNewProduct.ProductName.Like("%"+productName+"%"))
	}
	if recommendStatus, err := util.ParseInt32WithErr(recommendStatusStr); err != nil {
		conds = append(conds, smsHomeNewProduct.RecommendStatus.Eq(recommendStatus))
	}
	smsHomeNewProductDo := smsHomeNewProduct.Where(conds...)
	total, err := smsHomeNewProductDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsHomeNewProductDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}
