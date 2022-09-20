package service

import (
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/sms/service/dto"
	"mall-admin-server/util"
)

type SmsFlashPromotionProductRelationService struct {
	//
}

func (SmsFlashPromotionProductRelationService) Create(smsFlashPromotionProductRelation model.SmsFlashPromotionProductRelation) error {
	return query.SmsFlashPromotionProductRelation.Create(&smsFlashPromotionProductRelation)
}

func (SmsFlashPromotionProductRelationService) Update(idStr string, smsFlashPromotionProductRelation model.SmsFlashPromotionProductRelation) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsFlashPromotionProductRelation.Where(query.SmsFlashPromotionProductRelation.ID.Eq(id)).Updates(smsFlashPromotionProductRelation)
	return err
}

func (SmsFlashPromotionProductRelationService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsFlashPromotionProductRelation.Where(query.SmsFlashPromotionProductRelation.ID.Eq(id)).Delete()
	return err
}

func (SmsFlashPromotionProductRelationService) GetItem(idStr string) *model.SmsFlashPromotionProductRelation {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.SmsFlashPromotionProductRelation.Where(query.SmsFlashPromotionProductRelation.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (SmsFlashPromotionProductRelationService) List(flashPromotionIdStr, flashPromotionSessionIdStr, pageStr, sizeStr string) ([]*dto.SmsFlashPromotionProductRelationDto, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	flashPromotionId, err := util.ParseInt64WithErr(flashPromotionIdStr)
	if err != nil {
		return nil, 0
	}
	flashPromotionSessionId, err := util.ParseInt64WithErr(flashPromotionSessionIdStr)
	if err != nil {
		return nil, 0
	}
	total, err := query.SmsFlashPromotionProductRelation.GetListCount(flashPromotionId, flashPromotionSessionId)
	if err != nil {
		return nil, 0
	}
	data, err := query.SmsFlashPromotionProductRelation.GetListData(flashPromotionId, flashPromotionSessionId, offset, size)
	if err != nil {
		return nil, total
	}
	return data, total
}

func (SmsFlashPromotionProductRelationService) GetCount(flashPromotionIdStr, flashPromotionSessionIdStr string) int64 {
	flashPromotionId, err := util.ParseInt64WithErr(flashPromotionIdStr)
	if err != nil {
		return 0
	}
	flashPromotionSessionId, err := util.ParseInt64WithErr(flashPromotionSessionIdStr)
	if err != nil {
		return 0
	}
	smsFlashPromotionProductRelation := query.SmsFlashPromotionProductRelation
	total, err := smsFlashPromotionProductRelation.Where(
		smsFlashPromotionProductRelation.FlashPromotionID.Eq(flashPromotionId),
		smsFlashPromotionProductRelation.FlashPromotionSessionID.Eq(flashPromotionSessionId),
	).Count()
	if err != nil {
		return 0
	}
	return total
}
