package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/util"
)

// 首页专题推荐管理
type SmsHomeRecommendSubjectService struct {
	//
}

func (SmsHomeRecommendSubjectService) Create(smsHomeRecommendSubject model.SmsHomeRecommendSubject) error {
	smsHomeRecommendSubject.RecommendStatus = 1
	smsHomeRecommendSubject.Sort = 0
	return query.SmsHomeRecommendSubject.Create(&smsHomeRecommendSubject)
}

func (SmsHomeRecommendSubjectService) UpdateSort(idStr, sortStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	sort, err := util.ParseInt32WithErr(sortStr)
	if err != nil {
		return err
	}
	var smsHomeRecommendSubject model.SmsHomeRecommendSubject
	smsHomeRecommendSubject.Sort = sort
	_, err = query.SmsHomeRecommendSubject.Where(query.SmsHomeRecommendSubject.ID.Eq(id)).Updates(smsHomeRecommendSubject)
	return err
}

func (SmsHomeRecommendSubjectService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.SmsHomeRecommendSubject.Where(query.SmsHomeRecommendSubject.ID.In(ids...)).Delete()
	return err
}

func (SmsHomeRecommendSubjectService) UpdateRecommendStatus(idsStr []string, recommendStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	recommendStatus, err := util.ParseInt32WithErr(recommendStatusStr)
	if err != nil {
		return err
	}
	var smsHomeRecommendSubject model.SmsHomeRecommendSubject
	smsHomeRecommendSubject.RecommendStatus = recommendStatus
	_, err = query.SmsHomeRecommendSubject.Where(query.SmsHomeRecommendSubject.ID.In(ids...)).Updates(smsHomeRecommendSubject)
	return err
}

func (SmsHomeRecommendSubjectService) List(subjectName, recommendStatusStr, pageStr, sizeStr string) ([]*model.SmsHomeRecommendSubject, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsHomeRecommendSubject := query.SmsHomeRecommendSubject
	conds := make([]gen.Condition, 0)
	if subjectName != "" {
		conds = append(conds, smsHomeRecommendSubject.SubjectName.Like("%"+subjectName+"%"))
	}
	if recommendStatus, err := util.ParseInt32WithErr(recommendStatusStr); err != nil {
		conds = append(conds, smsHomeRecommendSubject.RecommendStatus.Eq(recommendStatus))
	}
	smsHomeRecommendSubjectDo := smsHomeRecommendSubject.Where(conds...)
	total, err := smsHomeRecommendSubjectDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsHomeRecommendSubjectDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}
