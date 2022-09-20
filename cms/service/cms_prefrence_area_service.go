package service

import (
	"mall-admin-server/cms/model"
	"mall-admin-server/cms/query"
)

// 商品优选管理
type CmsPrefrenceAreaService struct {
	//
}

func (CmsPrefrenceAreaService) ListAll() []*model.CmsPrefrenceArea {
	find, err := query.CmsPrefrenceArea.Find()
	if err != nil {
		return nil
	}
	return find
}
