package service

import (
	"gorm.io/gen"
	"mall-admin-server/cms/model"
	"mall-admin-server/cms/query"
	"mall-admin-server/util"
)

// 商品专题管理
type CmsSubjectService struct {
	//
}

func (CmsSubjectService) ListAll() []*model.CmsSubject {
	find, err := query.CmsSubject.Find()
	if err != nil {
		return nil
	}
	return find
}

func (CmsSubjectService) List(keyword, pageStr, sizeStr string) ([]*model.CmsSubject, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	cmsSubject := query.CmsSubject
	conds := make([]gen.Condition, 0)
	if keyword != "" {
		conds = append(conds, cmsSubject.Title.Like("%"+keyword+"%"))
	}
	cmsSubjectDo := cmsSubject.Where(conds...)
	count, err := cmsSubjectDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := cmsSubjectDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, count
	}
	return find, count
}
