// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"mall-admin-server/pms/model"
)

func newPmsProductFullReduction(db *gorm.DB) pmsProductFullReduction {
	_pmsProductFullReduction := pmsProductFullReduction{}

	_pmsProductFullReduction.pmsProductFullReductionDo.UseDB(db)
	_pmsProductFullReduction.pmsProductFullReductionDo.UseModel(&model.PmsProductFullReduction{})

	tableName := _pmsProductFullReduction.pmsProductFullReductionDo.TableName()
	_pmsProductFullReduction.ALL = field.NewField(tableName, "*")
	_pmsProductFullReduction.ID = field.NewInt64(tableName, "id")
	_pmsProductFullReduction.ProductID = field.NewInt64(tableName, "product_id")
	_pmsProductFullReduction.FullPrice = field.NewFloat64(tableName, "full_price")
	_pmsProductFullReduction.ReducePrice = field.NewFloat64(tableName, "reduce_price")

	_pmsProductFullReduction.fillFieldMap()

	return _pmsProductFullReduction
}

type pmsProductFullReduction struct {
	pmsProductFullReductionDo

	ALL         field.Field
	ID          field.Int64
	ProductID   field.Int64
	FullPrice   field.Float64
	ReducePrice field.Float64

	fieldMap map[string]field.Expr
}

func (p pmsProductFullReduction) Table(newTableName string) *pmsProductFullReduction {
	p.pmsProductFullReductionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductFullReduction) As(alias string) *pmsProductFullReduction {
	p.pmsProductFullReductionDo.DO = *(p.pmsProductFullReductionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductFullReduction) updateTableName(table string) *pmsProductFullReduction {
	p.ALL = field.NewField(table, "*")
	p.ID = field.NewInt64(table, "id")
	p.ProductID = field.NewInt64(table, "product_id")
	p.FullPrice = field.NewFloat64(table, "full_price")
	p.ReducePrice = field.NewFloat64(table, "reduce_price")

	p.fillFieldMap()

	return p
}

func (p *pmsProductFullReduction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductFullReduction) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["product_id"] = p.ProductID
	p.fieldMap["full_price"] = p.FullPrice
	p.fieldMap["reduce_price"] = p.ReducePrice
}

func (p pmsProductFullReduction) clone(db *gorm.DB) pmsProductFullReduction {
	p.pmsProductFullReductionDo.ReplaceDB(db)
	return p
}

type pmsProductFullReductionDo struct{ gen.DO }

func (p pmsProductFullReductionDo) Debug() *pmsProductFullReductionDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductFullReductionDo) WithContext(ctx context.Context) *pmsProductFullReductionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductFullReductionDo) ReadDB() *pmsProductFullReductionDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductFullReductionDo) WriteDB() *pmsProductFullReductionDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductFullReductionDo) Clauses(conds ...clause.Expression) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductFullReductionDo) Returning(value interface{}, columns ...string) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductFullReductionDo) Not(conds ...gen.Condition) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductFullReductionDo) Or(conds ...gen.Condition) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductFullReductionDo) Select(conds ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductFullReductionDo) Where(conds ...gen.Condition) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductFullReductionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pmsProductFullReductionDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pmsProductFullReductionDo) Order(conds ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductFullReductionDo) Distinct(cols ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductFullReductionDo) Omit(cols ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductFullReductionDo) Join(table schema.Tabler, on ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductFullReductionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductFullReductionDo) RightJoin(table schema.Tabler, on ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductFullReductionDo) Group(cols ...field.Expr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductFullReductionDo) Having(conds ...gen.Condition) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductFullReductionDo) Limit(limit int) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductFullReductionDo) Offset(offset int) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductFullReductionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductFullReductionDo) Unscoped() *pmsProductFullReductionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductFullReductionDo) Create(values ...*model.PmsProductFullReduction) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductFullReductionDo) CreateInBatches(values []*model.PmsProductFullReduction, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductFullReductionDo) Save(values ...*model.PmsProductFullReduction) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductFullReductionDo) First() (*model.PmsProductFullReduction, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductFullReduction), nil
	}
}

func (p pmsProductFullReductionDo) Take() (*model.PmsProductFullReduction, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductFullReduction), nil
	}
}

func (p pmsProductFullReductionDo) Last() (*model.PmsProductFullReduction, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductFullReduction), nil
	}
}

func (p pmsProductFullReductionDo) Find() ([]*model.PmsProductFullReduction, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductFullReduction), err
}

func (p pmsProductFullReductionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductFullReduction, err error) {
	buf := make([]*model.PmsProductFullReduction, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductFullReductionDo) FindInBatches(result *[]*model.PmsProductFullReduction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductFullReductionDo) Attrs(attrs ...field.AssignExpr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductFullReductionDo) Assign(attrs ...field.AssignExpr) *pmsProductFullReductionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductFullReductionDo) Joins(fields ...field.RelationField) *pmsProductFullReductionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductFullReductionDo) Preload(fields ...field.RelationField) *pmsProductFullReductionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductFullReductionDo) FirstOrInit() (*model.PmsProductFullReduction, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductFullReduction), nil
	}
}

func (p pmsProductFullReductionDo) FirstOrCreate() (*model.PmsProductFullReduction, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductFullReduction), nil
	}
}

func (p pmsProductFullReductionDo) FindByPage(offset int, limit int) (result []*model.PmsProductFullReduction, count int64, err error) {
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

func (p pmsProductFullReductionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductFullReductionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductFullReductionDo) Delete(models ...*model.PmsProductFullReduction) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductFullReductionDo) withDO(do gen.Dao) *pmsProductFullReductionDo {
	p.DO = *do.(*gen.DO)
	return p
}