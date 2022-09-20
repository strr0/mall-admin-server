package service

import (
	"gorm.io/gen"
	"mall-admin-server/oms/model"
	"mall-admin-server/oms/query"
	"mall-admin-server/oms/service/dto"
	"mall-admin-server/util"
	"time"
)

type OmsOrderService struct {
	//
}

func (OmsOrderService) List(omsOrderQueryDto dto.OmsOrderQueryDto, pageStr, sizeStr string) ([]*model.OmsOrder, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	omsOrder := query.OmsOrder
	conds := make([]gen.Condition, 0)
	conds = append(conds, omsOrder.DeleteStatus.Eq(0))
	if omsOrderQueryDto.OrderSn != "" {
		conds = append(conds, omsOrder.OrderSn.Eq(omsOrderQueryDto.OrderSn))
	}
	if omsOrderQueryDto.Status != 0 {
		conds = append(conds, omsOrder.Status.Eq(omsOrderQueryDto.Status))
	}
	if omsOrderQueryDto.SourceType != 0 {
		conds = append(conds, omsOrder.SourceType.Eq(omsOrderQueryDto.SourceType))
	}
	if omsOrderQueryDto.OrderType != 0 {
		conds = append(conds, omsOrder.OrderType.Eq(omsOrderQueryDto.OrderType))
	}
	if !omsOrderQueryDto.CreateTime.IsZero() {
		conds = append(conds, omsOrder.CreateTime.Gte(omsOrderQueryDto.CreateTime))
	}
	if omsOrderQueryDto.ReceiverKeyword != "" {
		conds = append(conds, omsOrder.Or(
			omsOrder.ReceiverName.Like("%"+omsOrderQueryDto.ReceiverKeyword+"%"),
			omsOrder.ReceiverPhone.Like("%"+omsOrderQueryDto.ReceiverKeyword+"%"),
		))
	}
	omsOrderDo := omsOrder.Where(conds...)
	count, err := omsOrderDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := omsOrderDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, count
	}
	return find, count
}

func (OmsOrderService) Delivery(deliveryList []dto.OmsOrderDeliveryDto) error {
	omsOrder := query.OmsOrder
	for _, deliver := range deliveryList {
		if deliver.OrderId != 0 {
			var order model.OmsOrder
			order.DeliverySn = deliver.DeliverySn
			order.DeliveryCompany = deliver.DeliveryCompany
			order.DeliveryTime = time.Now()
			order.Status = 2
			_, _ = omsOrder.Where(omsOrder.ID.Eq(deliver.OrderId), omsOrder.Status.Eq(1)).Updates(order)
		}
	}
	return nil
}

func (OmsOrderService) Clone(idsStr []string, note string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	var omsOrder model.OmsOrder
	omsOrder.Status = 4
	_, err := query.OmsOrder.Where(query.OmsOrder.DeleteStatus.Eq(0), query.OmsOrder.ID.In(ids...)).Updates(omsOrder)
	if err != nil {
		return err
	}
	historyList := make([]*model.OmsOrderOperateHistory, 0)
	for _, id := range ids {
		history := model.OmsOrderOperateHistory{
			OrderID:     id,
			CreateTime:  time.Now(),
			OperateMan:  "后台管理员",
			OrderStatus: 4,
			Note:        "订单关闭:" + note,
		}
		historyList = append(historyList, &history)
	}
	return query.OmsOrderOperateHistory.Create(historyList...)
}

func (OmsOrderService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.OmsOrder.Where(query.OmsOrder.ID.In(ids...)).Delete()
	return err
}

func (OmsOrderService) GetDetail(idStr string) *dto.OmsOrderDto {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.OmsOrder.Where(query.OmsOrder.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	orderItemList, _ := query.OmsOrderItem.Where(query.OmsOrderItem.OrderID.Eq(id)).Find()
	historyList, _ := query.OmsOrderOperateHistory.Where(query.OmsOrderOperateHistory.OrderID.Eq(id)).Find()
	return &dto.OmsOrderDto{
		OmsOrder:      *first,
		OrderItemList: orderItemList,
		HistoryList:   historyList,
	}
}

func (OmsOrderService) UpdateReceiverInfo(infoDto dto.OmsReceiverInfoDto) error {
	omsOrder := model.OmsOrder{
		ReceiverName:          infoDto.ReceiverName,
		ReceiverPhone:         infoDto.ReceiverPhone,
		ReceiverPostCode:      infoDto.ReceiverPostCode,
		ReceiverDetailAddress: infoDto.ReceiverDetailAddress,
		ReceiverProvince:      infoDto.ReceiverProvince,
		ReceiverCity:          infoDto.ReceiverCity,
		ReceiverRegion:        infoDto.ReceiverRegion,
		ModifyTime:            time.Now(),
	}
	_, err := query.OmsOrder.Where(query.OmsOrder.ID.Eq(infoDto.OrderId)).Updates(omsOrder)
	if err != nil {
		return err
	}
	history := model.OmsOrderOperateHistory{
		OrderID:     infoDto.OrderId,
		CreateTime:  time.Now(),
		OperateMan:  "后台管理员",
		OrderStatus: infoDto.Status,
		Note:        "修改收货人信息",
	}
	return query.OmsOrderOperateHistory.Create(&history)
}

func (OmsOrderService) UpdateMoneyInfo(infoDto dto.OmsMoneyInfoDto) error {
	omsOrder := model.OmsOrder{
		FreightAmount:  infoDto.FreightAmount,
		DiscountAmount: infoDto.DiscountAmount,
		ModifyTime:     time.Now(),
	}
	_, err := query.OmsOrder.Where(query.OmsOrder.ID.Eq(infoDto.OrderId)).Updates(omsOrder)
	if err != nil {
		return err
	}
	history := model.OmsOrderOperateHistory{
		CreateTime:  time.Now(),
		OperateMan:  "后台管理员",
		OrderStatus: infoDto.Status,
		Note:        "修改费用信息",
	}
	return query.OmsOrderOperateHistory.Create(&history)
}

func (OmsOrderService) UpdateNote(idStr, note, statusStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	omsOrder := model.OmsOrder{
		Note:       note,
		ModifyTime: time.Now(),
	}
	_, err = query.OmsOrder.Where(query.OmsOrder.ID.Eq(id)).Updates(omsOrder)
	if err != nil {
		return err
	}
	history := model.OmsOrderOperateHistory{
		OrderID:     id,
		CreateTime:  time.Now(),
		OperateMan:  "后台管理员",
		OrderStatus: status,
		Note:        "修改备注信息：" + note,
	}
	return query.OmsOrderOperateHistory.Create(&history)
}
