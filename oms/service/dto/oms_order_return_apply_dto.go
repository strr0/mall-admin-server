package dto

import "mall-admin-server/oms/model"

type OmsOrderReturnApplyDto struct {
	model.OmsOrderReturnApply
	model.OmsCompanyAddress
}
