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

	"mall-admin-server/oms/model"
)

func newOmsOrderReturnReason(db *gorm.DB) omsOrderReturnReason {
	_omsOrderReturnReason := omsOrderReturnReason{}

	_omsOrderReturnReason.omsOrderReturnReasonDo.UseDB(db)
	_omsOrderReturnReason.omsOrderReturnReasonDo.UseModel(&model.OmsOrderReturnReason{})

	tableName := _omsOrderReturnReason.omsOrderReturnReasonDo.TableName()
	_omsOrderReturnReason.ALL = field.NewField(tableName, "*")
	_omsOrderReturnReason.ID = field.NewInt64(tableName, "id")
	_omsOrderReturnReason.Name = field.NewString(tableName, "name")
	_omsOrderReturnReason.Sort = field.NewInt32(tableName, "sort")
	_omsOrderReturnReason.Status = field.NewInt32(tableName, "status")
	_omsOrderReturnReason.CreateTime = field.NewTime(tableName, "create_time")

	_omsOrderReturnReason.fillFieldMap()

	return _omsOrderReturnReason
}

type omsOrderReturnReason struct {
	omsOrderReturnReasonDo

	ALL        field.Field
	ID         field.Int64
	Name       field.String
	Sort       field.Int32
	Status     field.Int32
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (o omsOrderReturnReason) Table(newTableName string) *omsOrderReturnReason {
	o.omsOrderReturnReasonDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderReturnReason) As(alias string) *omsOrderReturnReason {
	o.omsOrderReturnReasonDo.DO = *(o.omsOrderReturnReasonDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderReturnReason) updateTableName(table string) *omsOrderReturnReason {
	o.ALL = field.NewField(table, "*")
	o.ID = field.NewInt64(table, "id")
	o.Name = field.NewString(table, "name")
	o.Sort = field.NewInt32(table, "sort")
	o.Status = field.NewInt32(table, "status")
	o.CreateTime = field.NewTime(table, "create_time")

	o.fillFieldMap()

	return o
}

func (o *omsOrderReturnReason) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderReturnReason) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 5)
	o.fieldMap["id"] = o.ID
	o.fieldMap["name"] = o.Name
	o.fieldMap["sort"] = o.Sort
	o.fieldMap["status"] = o.Status
	o.fieldMap["create_time"] = o.CreateTime
}

func (o omsOrderReturnReason) clone(db *gorm.DB) omsOrderReturnReason {
	o.omsOrderReturnReasonDo.ReplaceDB(db)
	return o
}

type omsOrderReturnReasonDo struct{ gen.DO }

func (o omsOrderReturnReasonDo) Debug() *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderReturnReasonDo) WithContext(ctx context.Context) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderReturnReasonDo) ReadDB() *omsOrderReturnReasonDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderReturnReasonDo) WriteDB() *omsOrderReturnReasonDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderReturnReasonDo) Clauses(conds ...clause.Expression) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderReturnReasonDo) Returning(value interface{}, columns ...string) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderReturnReasonDo) Not(conds ...gen.Condition) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderReturnReasonDo) Or(conds ...gen.Condition) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderReturnReasonDo) Select(conds ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderReturnReasonDo) Where(conds ...gen.Condition) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderReturnReasonDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *omsOrderReturnReasonDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o omsOrderReturnReasonDo) Order(conds ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderReturnReasonDo) Distinct(cols ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderReturnReasonDo) Omit(cols ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderReturnReasonDo) Join(table schema.Tabler, on ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderReturnReasonDo) LeftJoin(table schema.Tabler, on ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderReturnReasonDo) RightJoin(table schema.Tabler, on ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderReturnReasonDo) Group(cols ...field.Expr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderReturnReasonDo) Having(conds ...gen.Condition) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderReturnReasonDo) Limit(limit int) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderReturnReasonDo) Offset(offset int) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderReturnReasonDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderReturnReasonDo) Unscoped() *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderReturnReasonDo) Create(values ...*model.OmsOrderReturnReason) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderReturnReasonDo) CreateInBatches(values []*model.OmsOrderReturnReason, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderReturnReasonDo) Save(values ...*model.OmsOrderReturnReason) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderReturnReasonDo) First() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) Take() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) Last() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) Find() ([]*model.OmsOrderReturnReason, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderReturnReason), err
}

func (o omsOrderReturnReasonDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderReturnReason, err error) {
	buf := make([]*model.OmsOrderReturnReason, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderReturnReasonDo) FindInBatches(result *[]*model.OmsOrderReturnReason, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderReturnReasonDo) Attrs(attrs ...field.AssignExpr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderReturnReasonDo) Assign(attrs ...field.AssignExpr) *omsOrderReturnReasonDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderReturnReasonDo) Joins(fields ...field.RelationField) *omsOrderReturnReasonDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderReturnReasonDo) Preload(fields ...field.RelationField) *omsOrderReturnReasonDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderReturnReasonDo) FirstOrInit() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) FirstOrCreate() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) FindByPage(offset int, limit int) (result []*model.OmsOrderReturnReason, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o omsOrderReturnReasonDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderReturnReasonDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderReturnReasonDo) Delete(models ...*model.OmsOrderReturnReason) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderReturnReasonDo) withDO(do gen.Dao) *omsOrderReturnReasonDo {
	o.DO = *do.(*gen.DO)
	return o
}