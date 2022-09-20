package service

import (
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/query"
)

// 收货地址管理
type OmsCompanyAddressService struct {
	//
}

func (OmsCompanyAddressService) List() []*model.OmsCompanyAddress {
	find, err := query.OmsCompanyAddress.Find()
	if err != nil {
		return nil
	}
	return find
}
