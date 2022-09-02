package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Generator() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./ums/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db, _ := gorm.Open(mysql.Open("root:password@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModel("ums_admin"),
		g.GenerateModel("ums_member_level"),
		g.GenerateModel("ums_menu"),
		g.GenerateModel("ums_resource_category"),
		g.GenerateModel("ums_resource"),
		g.GenerateModel("ums_role"),
		g.GenerateModel("ums_admin_role_relation"),
		g.GenerateModel("ums_role_resource_relation"),
		g.GenerateModel("ums_role_menu_relation"),
	)
	g.Execute()
}
