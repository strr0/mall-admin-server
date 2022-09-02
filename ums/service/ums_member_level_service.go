package service

import (
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/query"
	"mall-admin-server/util"
)

type UmsMemberLevelService struct {
	//
}

func (UmsMemberLevelService) List(defaultStatusStr string) []*model.UmsMemberLevel {
	defaultStatus, err := util.ParseInt32WithErr(defaultStatusStr)
	if err != nil {
		return nil
	}
	find, err := query.UmsMemberLevel.Where(query.UmsMemberLevel.DefaultStatus.Eq(defaultStatus)).Find()
	if err != nil {
		return nil
	}
	return find
}
