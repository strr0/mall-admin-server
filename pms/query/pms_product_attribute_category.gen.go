// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service/dto"
)

func newPmsProductAttributeCategory(db *gorm.DB) pmsProductAttributeCategory {
	_pmsProductAttributeCategory := pmsProductAttributeCategory{}

	_pmsProductAttributeCategory.pmsProductAttributeCategoryDo.UseDB(db)
	_pmsProductAttributeCategory.pmsProductAttributeCategoryDo.UseModel(&model.PmsProductAttributeCategory{})

	tableName := _pmsProductAttributeCategory.pmsProductAttributeCategoryDo.TableName()
	_pmsProductAttributeCategory.ALL = field.NewField(tableName, "*")
	_pmsProductAttributeCategory.ID = field.NewInt64(tableName, "id")
	_pmsProductAttributeCategory.Name = field.NewString(tableName, "name")
	_pmsProductAttributeCategory.AttributeCount = field.NewInt32(tableName, "attribute_count")
	_pmsProductAttributeCategory.ParamCount = field.NewInt32(tableName, "param_count")

	_pmsProductAttributeCategory.fillFieldMap()

	return _pmsProductAttributeCategory
}

type pmsProductAttributeCategory struct {
	pmsProductAttributeCategoryDo

	ALL            field.Field
	ID             field.Int64
	Name           field.String
	AttributeCount field.Int32
	ParamCount     field.Int32

	fieldMap map[string]field.Expr
}

func (p pmsProductAttributeCategory) Table(newTableName string) *pmsProductAttributeCategory {
	p.pmsProductAttributeCategoryDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductAttributeCategory) As(alias string) *pmsProductAttributeCategory {
	p.pmsProductAttributeCategoryDo.DO = *(p.pmsProductAttributeCategoryDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductAttributeCategory) updateTableName(table string) *pmsProductAttributeCategory {
	p.ALL = field.NewField(table, "*")
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.AttributeCount = field.NewInt32(table, "attribute_count")
	p.ParamCount = field.NewInt32(table, "param_count")

	p.fillFieldMap()

	return p
}

func (p *pmsProductAttributeCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductAttributeCategory) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["attribute_count"] = p.AttributeCount
	p.fieldMap["param_count"] = p.ParamCount
}

func (p pmsProductAttributeCategory) clone(db *gorm.DB) pmsProductAttributeCategory {
	p.pmsProductAttributeCategoryDo.ReplaceDB(db)
	return p
}

type pmsProductAttributeCategoryDo struct{ gen.DO }

//sql(SELECT pac.id, pac.name, pa.id attr_id, pa.name attr_name
//FROM pms_product_attribute_category pac
//LEFT JOIN pms_product_attribute pa ON pac.id = pa.product_attribute_category_id AND pa.type=1)
func (p pmsProductAttributeCategoryDo) GetListWithAttr() (result []*dto.PmsProductAttributeCategoryDto, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT pac.id, pac.name, pa.id attr_id, pa.name attr_name FROM pms_product_attribute_category pac LEFT JOIN pms_product_attribute pa ON pac.id = pa.product_attribute_category_id AND pa.type=1 ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String()).Find(&result)
	err = executeSQL.Error
	return
}

func (p pmsProductAttributeCategoryDo) Debug() *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductAttributeCategoryDo) WithContext(ctx context.Context) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductAttributeCategoryDo) ReadDB() *pmsProductAttributeCategoryDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductAttributeCategoryDo) WriteDB() *pmsProductAttributeCategoryDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductAttributeCategoryDo) Clauses(conds ...clause.Expression) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductAttributeCategoryDo) Returning(value interface{}, columns ...string) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductAttributeCategoryDo) Not(conds ...gen.Condition) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductAttributeCategoryDo) Or(conds ...gen.Condition) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductAttributeCategoryDo) Select(conds ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductAttributeCategoryDo) Where(conds ...gen.Condition) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductAttributeCategoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pmsProductAttributeCategoryDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pmsProductAttributeCategoryDo) Order(conds ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductAttributeCategoryDo) Distinct(cols ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductAttributeCategoryDo) Omit(cols ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductAttributeCategoryDo) Join(table schema.Tabler, on ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductAttributeCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductAttributeCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductAttributeCategoryDo) Group(cols ...field.Expr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductAttributeCategoryDo) Having(conds ...gen.Condition) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductAttributeCategoryDo) Limit(limit int) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductAttributeCategoryDo) Offset(offset int) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductAttributeCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductAttributeCategoryDo) Unscoped() *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductAttributeCategoryDo) Create(values ...*model.PmsProductAttributeCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductAttributeCategoryDo) CreateInBatches(values []*model.PmsProductAttributeCategory, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductAttributeCategoryDo) Save(values ...*model.PmsProductAttributeCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductAttributeCategoryDo) First() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) Take() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) Last() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) Find() ([]*model.PmsProductAttributeCategory, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductAttributeCategory), err
}

func (p pmsProductAttributeCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductAttributeCategory, err error) {
	buf := make([]*model.PmsProductAttributeCategory, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductAttributeCategoryDo) FindInBatches(result *[]*model.PmsProductAttributeCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductAttributeCategoryDo) Attrs(attrs ...field.AssignExpr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductAttributeCategoryDo) Assign(attrs ...field.AssignExpr) *pmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductAttributeCategoryDo) Joins(fields ...field.RelationField) *pmsProductAttributeCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductAttributeCategoryDo) Preload(fields ...field.RelationField) *pmsProductAttributeCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductAttributeCategoryDo) FirstOrInit() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) FirstOrCreate() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) FindByPage(offset int, limit int) (result []*model.PmsProductAttributeCategory, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pmsProductAttributeCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductAttributeCategoryDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductAttributeCategoryDo) Delete(models ...*model.PmsProductAttributeCategory) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductAttributeCategoryDo) withDO(do gen.Dao) *pmsProductAttributeCategoryDo {
	p.DO = *do.(*gen.DO)
	return p
}
