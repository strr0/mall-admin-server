package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"mall-admin-server/oms/model"
)

func Generator() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./oms/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db, _ := gorm.Open(mysql.Open("root:password@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyInterface(func() {}, model.OmsCompanyAddress{})
	g.ApplyInterface(func() {}, model.OmsOrderReturnApply{})
	g.ApplyInterface(func() {}, model.OmsOrderReturnReason{})
	g.ApplyInterface(func() {}, model.OmsOrder{})
	g.ApplyInterface(func() {}, model.OmsOrderSetting{})
	g.ApplyBasic(
		g.GenerateModel("oms_order_item"),
		g.GenerateModel("oms_order_operate_history"),
	)
	g.Execute()
}
