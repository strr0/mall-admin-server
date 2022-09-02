package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"mall-admin-server/ums/model"
)

type SysRoleMethod interface {
	// sql(select r.* from ums_admin_role_relation ar left join ums_role r on ar.role_id = r.id where ar.admin_id = @adminId)
	GetRoleList(adminId int64) ([]*model.UmsRole, error)
}

type SysResourceMethod interface {
	// sql(SELECT ur.id id, ur.create_time, ur.`name`, ur.url, ur.description, ur.category_id
	// FROM ums_admin_role_relation ar
	// LEFT JOIN ums_role_resource_relation rr ON ar.role_id = rr.role_id
	// LEFT JOIN ums_resource ur ON ur.id = rr.resource_id
	// WHERE ar.admin_id = @adminId AND ur.id IS NOT NULL
	// GROUP BY ur.id)
	GetResourceList(adminId int64) ([]*model.UmsResource, error)

	// sql(SELECT r.id, r.create_time, r.`name`, r.url, r.description, r.category_id
	// FROM ums_role_resource_relation rrr
	// LEFT JOIN ums_resource r ON rrr.resource_id = r.id
	// WHERE rrr.role_id = @roleId AND r.id IS NOT NULL
	// GROUP BY r.id)
	GetResourceListByRoleId(roleId int64) ([]*model.UmsResource, error)
}

type SysMenuMethod interface {
	// sql(SELECT m.id id, m.parent_id, m.create_time, m.title, m.level, m.sort, m.name, m.icon, m.hidden
	// FROM ums_admin_role_relation arr
	// LEFT JOIN ums_role_menu_relation rmr ON arr.role_id = rmr.role_id
	// LEFT JOIN ums_menu m ON rmr.menu_id = m.id
	// WHERE arr.admin_id = @adminId AND m.id IS NOT NULL
	// GROUP BY m.id)
	GetMenuList(adminId int64) ([]*model.UmsMenu, error)

	// sql(SELECT m.id, m.parent_id, m.create_time, m.title, m.level, m.sort, m.name, m.icon, m.hidden
	// FROM ums_role_menu_relation rmr
	// LEFT JOIN ums_menu m ON rmr.menu_id = m.id
	// WHERE rmr.role_id = @roleId AND m.id IS NOT NULL
	// GROUP BY m.id
	GetMenuListByRoleId(roleId int64) ([]*model.UmsMenu, error)
}

func Generator() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./ums/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db, _ := gorm.Open(mysql.Open("root:password@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyInterface(func() {}, model.UmsAdmin{})
	g.ApplyInterface(func() {}, model.UmsMemberLevel{})
	g.ApplyInterface(func(method SysMenuMethod) {}, model.UmsMenu{})
	g.ApplyInterface(func() {}, model.UmsResourceCategory{})
	g.ApplyInterface(func(method SysResourceMethod) {}, model.UmsResource{})
	g.ApplyInterface(func(method SysRoleMethod) {}, model.UmsRole{})
	g.Execute()
}
