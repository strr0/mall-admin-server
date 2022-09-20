package service

import (
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/query"
	"mall-admin-server/util"
)

// 订单设置管理
type OmsOrderSettingService struct {
	//
}

func (OmsOrderSettingService) GetItem(idStr string) *model.OmsOrderSetting {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.OmsOrderSetting.Where(query.OmsOrderSetting.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (OmsOrderSettingService) Update(idStr string, omsOrderSetting model.OmsOrderSetting) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.OmsOrderSetting.Where(query.OmsOrderSetting.ID.Eq(id)).Updates(omsOrderSetting)
	return err
}
