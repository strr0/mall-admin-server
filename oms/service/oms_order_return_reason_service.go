package service

import (
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/query"
	"mall-admin-server/util"
	"time"
)

// 订单原因管理
type OmsOrderReturnReasonService struct {
	//
}

func (OmsOrderReturnReasonService) Create(omsOrderReturnReason model.OmsOrderReturnReason) error {
	omsOrderReturnReason.CreateTime = time.Now()
	return query.OmsOrderReturnReason.Create(&omsOrderReturnReason)
}

func (OmsOrderReturnReasonService) Update(idStr string, omsOrderReturnReason model.OmsOrderReturnReason) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.OmsOrderReturnReason.Where(query.OmsOrderReturnReason.ID.Eq(id)).Updates(omsOrderReturnReason)
	return err
}

func (OmsOrderReturnReasonService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.OmsOrderReturnReason.Where(query.OmsOrderReturnReason.ID.In(ids...)).Delete()
	return err
}

func (OmsOrderReturnReasonService) List(pageStr, sizeStr string) ([]*model.OmsOrderReturnReason, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	count, err := query.OmsOrderReturnReason.Count()
	if err != nil {
		return nil, 0
	}
	find, err := query.OmsOrderReturnReason.Order(query.OmsOrderReturnReason.Sort.Desc()).Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, count
	}
	return find, count
}

func (OmsOrderReturnReasonService) UpdateStatus(idsStr []string, statusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	_, err = query.OmsOrderReturnReason.Where(query.OmsOrderReturnReason.ID.In(ids...)).Update(query.OmsOrderReturnReason.Status, status)
	return err
}

func (OmsOrderReturnReasonService) GetItem(idStr string) *model.OmsOrderReturnReason {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.OmsOrderReturnReason.Where(query.OmsOrderReturnReason.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}
