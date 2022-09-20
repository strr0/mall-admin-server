package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/util"
)

// 首页品牌管理
type SmsHomeBrandService struct {
	//
}

func (SmsHomeBrandService) Create(smsHomeBrand model.SmsHomeBrand) error {
	smsHomeBrand.RecommendStatus = 1
	smsHomeBrand.Sort = 0
	return query.SmsHomeBrand.Create(&smsHomeBrand)
}

func (SmsHomeBrandService) UpdateSort(idStr, sortStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	sort, err := util.ParseInt32WithErr(sortStr)
	if err != nil {
		return err
	}
	var smsHomeBrand model.SmsHomeBrand
	smsHomeBrand.Sort = sort
	_, err = query.SmsHomeBrand.Where(query.SmsHomeBrand.ID.Eq(id)).Updates(smsHomeBrand)
	return err
}

func (SmsHomeBrandService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.SmsHomeBrand.Where(query.SmsHomeBrand.ID.In(ids...)).Delete()
	return err
}

func (SmsHomeBrandService) UpdateRecommendStatus(idsStr []string, recommendStatusStr string) error {
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
	var smsHomeBrand model.SmsHomeBrand
	smsHomeBrand.RecommendStatus = recommendStatus
	_, err = query.SmsHomeBrand.Where(query.SmsHomeBrand.ID.In(ids...)).Updates(smsHomeBrand)
	return err
}

func (SmsHomeBrandService) List(brandName, recommendStatusStr, pageStr, sizeStr string) ([]*model.SmsHomeBrand, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsHomeBrand := query.SmsHomeBrand
	conds := make([]gen.Condition, 0)
	if brandName != "" {
		conds = append(conds, smsHomeBrand.BrandName.Like("%"+brandName+"%"))
	}
	if recommendStatus, err := util.ParseInt32WithErr(recommendStatusStr); err != nil {
		conds = append(conds, smsHomeBrand.RecommendStatus.Eq(recommendStatus))
	}
	smsHomeBrandDo := smsHomeBrand.Where(conds...)
	total, err := smsHomeBrandDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsHomeBrandDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}
