package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/util"
)

// 首页人气推荐管理
type SmsHomeRecommendProductService struct {
	//
}

func (SmsHomeRecommendProductService) Create(smsHomeRecommendProduct model.SmsHomeRecommendProduct) error {
	smsHomeRecommendProduct.RecommendStatus = 1
	smsHomeRecommendProduct.Sort = 0
	return query.SmsHomeRecommendProduct.Create(&smsHomeRecommendProduct)
}

func (SmsHomeRecommendProductService) UpdateSort(idStr, sortStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	sort, err := util.ParseInt32WithErr(sortStr)
	if err != nil {
		return err
	}
	var smsHomeRecommendProduct model.SmsHomeRecommendProduct
	smsHomeRecommendProduct.Sort = sort
	_, err = query.SmsHomeRecommendProduct.Where(query.SmsHomeRecommendProduct.ID.Eq(id)).Updates(smsHomeRecommendProduct)
	return err
}

func (SmsHomeRecommendProductService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.SmsHomeRecommendProduct.Where(query.SmsHomeRecommendProduct.ID.In(ids...)).Delete()
	return err
}

func (SmsHomeRecommendProductService) UpdateRecommendStatus(idsStr []string, recommendStatusStr string) error {
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
	var smsHomeRecommendProduct model.SmsHomeRecommendProduct
	smsHomeRecommendProduct.RecommendStatus = recommendStatus
	_, err = query.SmsHomeRecommendProduct.Where(query.SmsHomeRecommendProduct.ID.In(ids...)).Updates(smsHomeRecommendProduct)
	return err
}

func (SmsHomeRecommendProductService) List(productName, recommendStatusStr, pageStr, sizeStr string) ([]*model.SmsHomeRecommendProduct, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsHomeRecommendProduct := query.SmsHomeRecommendProduct
	conds := make([]gen.Condition, 0)
	if productName != "" {
		conds = append(conds, smsHomeRecommendProduct.ProductName.Like("%"+productName+"%"))
	}
	if recommendStatus, err := util.ParseInt32WithErr(recommendStatusStr); err != nil {
		conds = append(conds, smsHomeRecommendProduct.RecommendStatus.Eq(recommendStatus))
	}
	smsHomeRecommendProductDo := smsHomeRecommendProduct.Where(conds...)
	total, err := smsHomeRecommendProductDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsHomeRecommendProductDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}
