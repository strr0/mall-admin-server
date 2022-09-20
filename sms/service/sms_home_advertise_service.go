package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/util"
)

// 首页广告
type SmsHomeAdvertiseService struct {
	//
}

func (SmsHomeAdvertiseService) Create(smsHomeAdvertise model.SmsHomeAdvertise) error {
	smsHomeAdvertise.ClickCount = 0
	smsHomeAdvertise.OrderCount = 0
	return query.SmsHomeAdvertise.Create(&smsHomeAdvertise)
}

func (SmsHomeAdvertiseService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.SmsHomeAdvertise.Where(query.SmsHomeAdvertise.ID.In(ids...)).Delete()
	return err
}

func (SmsHomeAdvertiseService) UpdateStatus(idStr, statusStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	var smsHomeAdvertise model.SmsHomeAdvertise
	smsHomeAdvertise.Status = status
	_, err = query.SmsHomeAdvertise.Where(query.SmsHomeAdvertise.ID.Eq(id)).Updates(smsHomeAdvertise)
	return err
}

func (SmsHomeAdvertiseService) GetItem(idStr string) *model.SmsHomeAdvertise {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.SmsHomeAdvertise.Where(query.SmsHomeAdvertise.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (SmsHomeAdvertiseService) Update(idStr string, smsHomeAdvertise model.SmsHomeAdvertise) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsHomeAdvertise.Where(query.SmsHomeAdvertise.ID.Eq(id)).Updates(smsHomeAdvertise)
	return err
}

func (SmsHomeAdvertiseService) List(name, typeStr, endTimeStr, pageStr, sizeStr string) ([]*model.SmsHomeAdvertise, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsHomeAdvertise := query.SmsHomeAdvertise
	conds := make([]gen.Condition, 0)
	if name != "" {
		conds = append(conds, smsHomeAdvertise.Name.Like("%"+name+"%"))
	}
	if type_, err := util.ParseInt32WithErr(typeStr); err != nil {
		conds = append(conds, smsHomeAdvertise.Type.Eq(type_))
	}
	if endTimeStr != "" {
		startTime, err := util.ParseTimeWithErr(endTimeStr + " 00:00:00")
		if err != nil {
			return nil, 0
		}
		endTime, err := util.ParseTimeWithErr(endTimeStr + " 23:59:59")
		if err != nil {
			return nil, 0
		}
		conds = append(conds, smsHomeAdvertise.EndTime.Between(startTime, endTime))
	}
	smsHomeAdvertiseDo := smsHomeAdvertise.Where(conds...)
	total, err := smsHomeAdvertiseDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsHomeAdvertiseDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}
