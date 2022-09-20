package service

import (
	"gorm.io/gen"
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/query"
	"mall-admin-server/oms/service/dto"
	"mall-admin-server/util"
	"time"
)

// 订单退货管理
type OmsOrderReturnApplyService struct {
	//
}

func (OmsOrderReturnApplyService) List(queryDto dto.OmsOrderReturnApplyQueryDto, pageStr, sizeStr string) ([]*model.OmsOrderReturnApply, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	omsOrderReturnApply := query.OmsOrderReturnApply
	conds := make([]gen.Condition, 0)
	if queryDto.Id != 0 {
		conds = append(conds, omsOrderReturnApply.ID.Eq(queryDto.Id))
	}
	if queryDto.Status != 0 {
		conds = append(conds, omsOrderReturnApply.Status.Eq(queryDto.Status))
	}
	if queryDto.HandleMan != "" {
		conds = append(conds, omsOrderReturnApply.HandleMan.Eq(queryDto.HandleMan))
	}
	if !queryDto.CreateTime.IsZero() {
		conds = append(conds, omsOrderReturnApply.CreateTime.Gte(queryDto.CreateTime))
	}
	if !queryDto.HandleTime.IsZero() {
		conds = append(conds, omsOrderReturnApply.HandleTime.Gte(queryDto.HandleTime))
	}
	if queryDto.ReceiverKeyword != "" {
		conds = append(conds, omsOrderReturnApply.Or(
			omsOrderReturnApply.ReturnName.Like("%"+queryDto.ReceiverKeyword+"%"),
			omsOrderReturnApply.ReturnPhone.Like("%"+queryDto.ReceiverKeyword+"%"),
		))
	}
	omsOrderReturnApplyDo := omsOrderReturnApply.Where(conds...)
	count, err := omsOrderReturnApplyDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := omsOrderReturnApplyDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, count
	}
	return find, count
}

func (OmsOrderReturnApplyService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.OmsOrderReturnApply.Where(query.OmsOrderReturnApply.ID.In(ids...)).Delete()
	return err
}

func (OmsOrderReturnApplyService) UpdateStatus(idStr string, statusDto dto.OmsUpdateStatusDto) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
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
	_, err = query.OmsOrderReturnApply.Where(query.OmsOrderReturnApply.ID.Eq(id)).Updates(returnApply)
	return err
}

func (OmsOrderReturnApplyService) GetItem(idStr string) *dto.OmsOrderReturnApplyDto {
	apply := query.OmsOrderReturnApply
	address := query.OmsCompanyAddress
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	var omsOrderReturnApplyDto dto.OmsOrderReturnApplyDto
	err = apply.LeftJoin(address, address.ID.EqCol(apply.CompanyAddressID)).Where(apply.ID.Eq(id)).Scan(&omsOrderReturnApplyDto)
	if err != nil {
		return nil
	}
	return &omsOrderReturnApplyDto
}
