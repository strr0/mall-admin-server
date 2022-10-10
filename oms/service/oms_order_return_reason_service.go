package service

import (
	"gorm.io/gorm"
	"mall-admin-server/oms/model"
	"mall-admin-server/util"
	"time"
)

// 订单原因管理
type OmsOrderReturnReasonService struct {
	DB *gorm.DB
}

func (iService OmsOrderReturnReasonService) Create(omsOrderReturnReason model.OmsOrderReturnReason) error {
	omsOrderReturnReason.CreateTime = time.Now()
	result := iService.DB.Create(&omsOrderReturnReason)
	return result.Error
}

func (iService OmsOrderReturnReasonService) Update(idStr string, omsOrderReturnReason model.OmsOrderReturnReason) error {
	result := iService.DB.Save(&omsOrderReturnReason)
	return result.Error
}

func (iService OmsOrderReturnReasonService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.OmsOrderReturnReason{}, ids)
	return result.Error
}

func (iService OmsOrderReturnReasonService) List(pageStr, sizeStr string) ([]model.OmsOrderReturnReason, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	var count int64
	result := iService.DB.Model(&model.OmsOrderReturnReason{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.OmsOrderReturnReason
	result = iService.DB.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService OmsOrderReturnReasonService) UpdateStatus(ids []string, status string) error {
	result := iService.DB.Model(&model.OmsOrderReturnReason{}).Where("id in ?", ids).Update("status", status)
	return result.Error
}

func (iService OmsOrderReturnReasonService) GetItem(id string) *model.OmsOrderReturnReason {
	var omsOrderReturnReason model.OmsOrderReturnReason
	result := iService.DB.First(&omsOrderReturnReason, id)
	if result.Error != nil {
		return nil
	}
	return &omsOrderReturnReason
}
