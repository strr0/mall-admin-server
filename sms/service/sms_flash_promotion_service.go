package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/util"
)

// 限购活动
type SmsFlashPromotionService struct {
	//
}

func (SmsFlashPromotionService) Create(smsFlashPromotion model.SmsFlashPromotion) error {
	return query.SmsFlashPromotion.Create(&smsFlashPromotion)
}

func (SmsFlashPromotionService) Update(idStr string, smsFlashPromotion model.SmsFlashPromotion) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsFlashPromotion.Where(query.SmsFlashPromotion.ID.Eq(id)).Updates(smsFlashPromotion)
	return err
}

func (SmsFlashPromotionService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsFlashPromotion.Where(query.SmsFlashPromotion.ID.Eq(id)).Delete()
	return err
}

func (SmsFlashPromotionService) UpdateStatus(idStr, statusStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	var smsFlashPromotion model.SmsFlashPromotion
	smsFlashPromotion.Status = status
	_, err = query.SmsFlashPromotion.Where(query.SmsFlashPromotion.ID.Eq(id)).Updates(smsFlashPromotion)
	return err
}

func (SmsFlashPromotionService) GetItem(idStr string) *model.SmsFlashPromotion {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.SmsFlashPromotion.Where(query.SmsFlashPromotion.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (SmsFlashPromotionService) List(keyword, pageStr, sizeStr string) ([]*model.SmsFlashPromotion, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsFlashPromotion := query.SmsFlashPromotion
	conds := make([]gen.Condition, 0)
	if keyword != "" {
		conds = append(conds, smsFlashPromotion.Title.Like("%"+keyword+"%"))
	}
	smsFlashPromotion.Where(conds...)
	total, err := smsFlashPromotion.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsFlashPromotion.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}
