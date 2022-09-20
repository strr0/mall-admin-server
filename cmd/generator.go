package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Generator() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pms/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db, _ := gorm.Open(mysql.Open("root:password@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModel("pms_brand"),
		g.GenerateModel("pms_product_attribute_category"),
		g.GenerateModel("pms_product_attribute"),
		g.GenerateModel("pms_product_category"),
		g.GenerateModel("pms_product"),
		g.GenerateModel("pms_sku_stock"),
	)
	g.Execute()
}
