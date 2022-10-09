package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/util"
)

// 首页专题推荐管理
type SmsHomeRecommendSubjectService struct {
	DB *gorm.DB
}

func (iService SmsHomeRecommendSubjectService) Create(smsHomeRecommendSubject model.SmsHomeRecommendSubject) error {
	smsHomeRecommendSubject.RecommendStatus = 1
	smsHomeRecommendSubject.Sort = 0
	result := iService.DB.Create(&smsHomeRecommendSubject)
	return result.Error
}

func (iService SmsHomeRecommendSubjectService) UpdateSort(id, sort string) error {
	result := iService.DB.Model(&model.SmsHomeRecommendSubject{}).Where("id = ?", id).Update("sort", sort)
	return result.Error
}

func (iService SmsHomeRecommendSubjectService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.SmsHomeRecommendSubject{}, ids)
	return result.Error
}

func (iService SmsHomeRecommendSubjectService) UpdateRecommendStatus(ids []string, recommendStatus string) error {
	result := iService.DB.Model(&model.SmsHomeRecommendSubject{}).Where("id in ?", ids).Update("recommend_status", recommendStatus)
	return result.Error
}

func (iService SmsHomeRecommendSubjectService) List(subjectName, recommendStatus, pageStr, sizeStr string) ([]model.SmsHomeRecommendSubject, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if subjectName != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("subject_name like ?", "%"+subjectName+"%")
		})
	}
	if recommendStatus != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("recommend_status = ?", recommendStatus)
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.SmsHomeRecommendSubject{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.SmsHomeRecommendSubject
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
