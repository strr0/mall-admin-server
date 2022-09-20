package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service/dto"
)

type PmsProductAttributeCategoryMethod interface {
	// sql(SELECT pac.id, pac.name, pa.id attr_id, pa.name attr_name
	// FROM pms_product_attribute_category pac
	// LEFT JOIN pms_product_attribute pa ON pac.id = pa.product_attribute_category_id AND pa.type=1)
	GetListWithAttr() ([]*dto.PmsProductAttributeCategoryDto, error)
}

type PmsProductAttributeMethod interface {
	// sql(SELECT pa.id  attribute_id, pac.id attribute_category_id
	// FROM pms_product_category_attribute_relation pcar
	// LEFT JOIN pms_product_attribute pa ON pa.id = pcar.product_attribute_id
	// LEFT JOIN pms_product_attribute_category pac ON pa.product_attribute_category_id = pac.id
	// WHERE pcar.product_category_id = @id)
	GetProductAttrInfo(id int64) ([]*dto.PmsProductAttributeDto, error)
}

func Generator() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pms/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db, _ := gorm.Open(mysql.Open("root:password@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyInterface(func() {}, model.PmsBrand{})
	g.ApplyInterface(func(method PmsProductAttributeCategoryMethod) {}, model.PmsProductAttributeCategory{})
	g.ApplyInterface(func(method PmsProductAttributeMethod) {}, model.PmsProductAttribute{})
	g.ApplyInterface(func() {}, model.PmsProductCategoryAttributeRelation{})
	g.ApplyInterface(func() {}, model.PmsProductCategory{})
	g.ApplyInterface(func() {}, model.PmsProduct{})
	g.ApplyInterface(func() {}, model.PmsSkuStock{})
	g.ApplyInterface(func() {}, model.PmsProductLadder{})
	g.ApplyInterface(func() {}, model.PmsProductFullReduction{})
	g.ApplyInterface(func() {}, model.PmsMemberPrice{})
	g.ApplyInterface(func() {}, model.PmsProductAttributeValue{})
	g.ApplyBasic(
		g.GenerateModel("pms_product_vertify_record"),
	)
	g.Execute()
}
