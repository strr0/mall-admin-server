package service

import (
	"gorm.io/gorm"
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/service/dto"
	"mall-admin-server/util"
	"time"
)

type OmsOrderService struct {
	DB *gorm.DB
}

func (iService OmsOrderService) List(omsOrderQueryDto dto.OmsOrderQueryDto, pageStr, sizeStr string) ([]model.OmsOrder, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where("delete_status = ?", 0)
	})
	if omsOrderQueryDto.OrderSn != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("order_sn = ?", omsOrderQueryDto.OrderSn)
		})
	}
	if omsOrderQueryDto.Status != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", omsOrderQueryDto.Status)
		})
	}
	if omsOrderQueryDto.SourceType != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("source_type = ?", omsOrderQueryDto.SourceType)
		})
	}
	if omsOrderQueryDto.OrderType != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("order_type = ?", omsOrderQueryDto.OrderType)
		})
	}
	if !omsOrderQueryDto.CreateTime.IsZero() {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("create_time >= ?", omsOrderQueryDto.CreateTime)
		})
	}
	if omsOrderQueryDto.ReceiverKeyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("(receiver_name like ? or receiver_phone like ?)", "%"+omsOrderQueryDto.ReceiverKeyword+"%", "%"+omsOrderQueryDto.ReceiverKeyword+"%")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.OmsOrder{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.OmsOrder
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService OmsOrderService) Delivery(deliveryList []dto.OmsOrderDeliveryDto) error {
	for _, deliver := range deliveryList {
		if deliver.OrderId != 0 {
			var order model.OmsOrder
			order.DeliverySn = deliver.DeliverySn
			order.DeliveryCompany = deliver.DeliveryCompany
			order.DeliveryTime = time.Now()
			order.Status = 2
			iService.DB.Model(&model.OmsOrder{}).Where("order_id = ? and status = ?", deliver.OrderId, 1).Updates(order)
		}
	}
	return nil
}

func (iService OmsOrderService) Close(idsStr []string, note string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	result := iService.DB.Model(&model.OmsOrder{}).Where("delete_status = ? and id in ?", 0, ids).Update("status", 4)
	if result.Error != nil {
		return result.Error
	}
	historyList := make([]model.OmsOrderOperateHistory, 0)
	for _, id := range ids {
		history := model.OmsOrderOperateHistory{
			OrderID:     id,
			CreateTime:  time.Now(),
			OperateMan:  "后台管理员",
			OrderStatus: 4,
			Note:        "订单关闭:" + note,
		}
		historyList = append(historyList, history)
	}
	result = iService.DB.Create(&historyList)
	return result.Error
}

func (iService OmsOrderService) Delete(ids []string) error {
	result := iService.DB.Delete(&model.OmsOrder{}, ids)
	return result.Error
}

func (iService OmsOrderService) GetDetail(id string) *dto.OmsOrderDto {
	var omsOrder model.OmsOrder
	result := iService.DB.First(&omsOrder, id)
	if result.Error != nil {
		return nil
	}
	var orderItemList []model.OmsOrderItem
	var historyList []model.OmsOrderOperateHistory
	result = iService.DB.Where("order_id = ?", id).Find(&orderItemList)
	result = iService.DB.Where("order_id = ?", id).Find(&historyList)
	return &dto.OmsOrderDto{
		OmsOrder:      omsOrder,
		OrderItemList: orderItemList,
		HistoryList:   historyList,
	}
}

func (iService OmsOrderService) UpdateReceiverInfo(infoDto dto.OmsReceiverInfoDto) error {
	omsOrder := model.OmsOrder{
		ID:                    infoDto.OrderId,
		ReceiverName:          infoDto.ReceiverName,
		ReceiverPhone:         infoDto.ReceiverPhone,
		ReceiverPostCode:      infoDto.ReceiverPostCode,
		ReceiverDetailAddress: infoDto.ReceiverDetailAddress,
		ReceiverProvince:      infoDto.ReceiverProvince,
		ReceiverCity:          infoDto.ReceiverCity,
		ReceiverRegion:        infoDto.ReceiverRegion,
		ModifyTime:            time.Now(),
	}
	result := iService.DB.Model(&omsOrder).Updates(omsOrder)
	if result.Error != nil {
		return result.Error
	}
	history := model.OmsOrderOperateHistory{
		OrderID:     infoDto.OrderId,
		CreateTime:  time.Now(),
		OperateMan:  "后台管理员",
		OrderStatus: infoDto.Status,
		Note:        "修改收货人信息",
	}
	result = iService.DB.Create(&history)
	return result.Error
}

func (iService OmsOrderService) UpdateMoneyInfo(infoDto dto.OmsMoneyInfoDto) error {
	omsOrder := model.OmsOrder{
		ID:             infoDto.OrderId,
		FreightAmount:  infoDto.FreightAmount,
		DiscountAmount: infoDto.DiscountAmount,
		ModifyTime:     time.Now(),
	}
	result := iService.DB.Model(&omsOrder).Updates(omsOrder)
	if result.Error != nil {
		return result.Error
	}
	history := model.OmsOrderOperateHistory{
		CreateTime:  time.Now(),
		OperateMan:  "后台管理员",
		OrderStatus: infoDto.Status,
		Note:        "修改费用信息",
	}
	result = iService.DB.Create(&history)
	return result.Error
}

func (iService OmsOrderService) UpdateNote(idStr, note, statusStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	omsOrder := model.OmsOrder{
		ID:         id,
		Note:       note,
		ModifyTime: time.Now(),
	}
	result := iService.DB.Model(&omsOrder).Updates(omsOrder)
	if result.Error != nil {
		return err
	}
	history := model.OmsOrderOperateHistory{
		OrderID:     id,
		CreateTime:  time.Now(),
		OperateMan:  "后台管理员",
		OrderStatus: status,
		Note:        "修改备注信息：" + note,
	}
	result = iService.DB.Create(&history)
	return result.Error
}
