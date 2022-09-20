package service

import (
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/sms/service/dto"
	"mall-admin-server/util"
	"time"
)

// 限购场次
type SmsFlashPromotionSessionService struct {
	//
}

func (SmsFlashPromotionSessionService) Create(smsFlashPromotionSession model.SmsFlashPromotionSession) error {
	smsFlashPromotionSession.CreateTime = time.Now()
	return query.SmsFlashPromotionSession.Create(&smsFlashPromotionSession)
}

func (SmsFlashPromotionSessionService) Update(idStr string, smsFlashPromotionSession model.SmsFlashPromotionSession) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsFlashPromotionSession.Where(query.SmsFlashPromotionSession.ID.Eq(id)).Updates(smsFlashPromotionSession)
	return err
}

func (SmsFlashPromotionSessionService) UpdateStatus(idStr, statusStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	var smsFlashPromotionSession model.SmsFlashPromotionSession
	smsFlashPromotionSession.Status = status
	_, err = query.SmsFlashPromotionSession.Where(query.SmsFlashPromotionSession.ID.Eq(id)).Updates(smsFlashPromotionSession)
	return err
}

func (SmsFlashPromotionSessionService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsFlashPromotionSession.Where(query.SmsFlashPromotionSession.ID.Eq(id)).Delete()
	return err
}

func (SmsFlashPromotionSessionService) GetItem(idStr string) *model.SmsFlashPromotionSession {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.SmsFlashPromotionSession.Where(query.SmsFlashPromotionSession.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (SmsFlashPromotionSessionService) ListAll() []*model.SmsFlashPromotionSession {
	find, err := query.SmsFlashPromotionSession.Find()
	if err != nil {
		return nil
	}
	return find
}

func (SmsFlashPromotionSessionService) SelectList(flashPromotionIdStr string) []*dto.SmsFlashPromotionSessionDto {
	flashPromotionId, err := util.ParseInt64WithErr(flashPromotionIdStr)
	if err != nil {
		return nil
	}
	result := make([]*dto.SmsFlashPromotionSessionDto, 0)
	find, err := query.SmsFlashPromotionSession.Where(query.SmsFlashPromotionSession.Status.Eq(1)).Find()
	if err != nil {
		return nil
	}
	for _, flashPromotionSession := range find {
		count, err := query.SmsFlashPromotionProductRelation.Where(
			query.SmsFlashPromotionProductRelation.FlashPromotionID.Eq(flashPromotionId),
			query.SmsFlashPromotionProductRelation.FlashPromotionSessionID.Eq(flashPromotionSession.ID),
		).Count()
		if err == nil {
			result = append(result, &dto.SmsFlashPromotionSessionDto{
				SmsFlashPromotionSession: *flashPromotionSession,
				ProductCount:             count,
			})
		}
	}
	return result
}
