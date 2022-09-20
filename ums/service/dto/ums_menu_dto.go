package dto

import "mall-admin-server/ums/model"

type UmsMenuDto struct {
	model.UmsMenu
	Children []UmsMenuDto
}
