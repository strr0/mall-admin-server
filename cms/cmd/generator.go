package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Generator() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./cms/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db, _ := gorm.Open(mysql.Open("root:password@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModel("cms_prefrence_area"),
		g.GenerateModel("cms_subject"),
		g.GenerateModel("cms_prefrence_area_product_relation"),
		g.GenerateModel("cms_subject_product_relation"),
	)
	g.Execute()
}
