package service

import (
	"gorm.io/gorm"
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/service/dto"
	"mall-admin-server/util"
	"time"
)

// 订单退货管理
type OmsOrderReturnApplyService struct {
	DB *gorm.DB
}

func (iService OmsOrderReturnApplyService) List(queryDto dto.OmsOrderReturnApplyQueryDto, pageStr, sizeStr string) ([]model.OmsOrderReturnApply, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if queryDto.Id != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", queryDto.Id)
		})
	}
	if queryDto.Status != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", queryDto.Status)
		})
	}
	if queryDto.HandleMan != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("handle_man = ?", queryDto.HandleMan)
		})
	}
	if !queryDto.CreateTime.IsZero() {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("create_time >= ?", queryDto.CreateTime)
		})
	}
	if !queryDto.HandleTime.IsZero() {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("handle_time >= ?", queryDto.HandleTime)
		})
	}
	if queryDto.ReceiverKeyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("(return_name like ? or return_phone like ?)", "%"+queryDto.ReceiverKeyword+"%", "%"+queryDto.ReceiverKeyword+"%")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.OmsOrderReturnApply{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.OmsOrderReturnApply
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService OmsOrderReturnApplyService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.OmsOrderReturnApply{}, ids)
	return result.Error
}

func (iService OmsOrderReturnApplyService) UpdateStatus(idStr string, statusDto dto.OmsUpdateStatusDto) error {
	var returnApply model.OmsOrderReturnApply
	switch statusDto.Status {
	case 1:
		returnApply.Status = statusDto.Status
		returnApply.ReturnAmount = statusDto.ReturnAmount
		returnApply.CompanyAddressID = statusDto.CompanyAddressId
		returnApply.HandleTime = time.Now()
		returnApply.HandleMan = statusDto.HandleMan
		returnApply.HandleNote = statusDto.HandleNote
	case 2:
		returnApply.Status = 2
		returnApply.ReceiveTime = time.Now()
		returnApply.ReceiveMan = statusDto.ReceiveMan
		returnApply.HandleNote = statusDto.HandleNote
	case 3:
		returnApply.Status = 3
		returnApply.HandleTime = time.Now()
		returnApply.HandleMan = statusDto.HandleMan
		returnApply.HandleNote = statusDto.HandleNote
	}
	result := iService.DB.Save(&returnApply)
	return result.Error
}

func (iService OmsOrderReturnApplyService) GetItem(id string) *dto.OmsOrderReturnApplyDto {
	var omsOrderReturnApplyDto dto.OmsOrderReturnApplyDto
	result := iService.DB.Table("oms_order_return_apply t1").Select("t1.*, t2.*").Joins("left join oms_company_address t2 on t1.company_address_id = t2.id").Where("t1.apply_id = ?", id).Scan(&omsOrderReturnApplyDto)
	if result.Error != nil {
		return nil
	}
	return &omsOrderReturnApplyDto
}
