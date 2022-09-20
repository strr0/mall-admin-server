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

func newOmsOrderSetting(db *gorm.DB) omsOrderSetting {
	_omsOrderSetting := omsOrderSetting{}

	_omsOrderSetting.omsOrderSettingDo.UseDB(db)
	_omsOrderSetting.omsOrderSettingDo.UseModel(&model.OmsOrderSetting{})

	tableName := _omsOrderSetting.omsOrderSettingDo.TableName()
	_omsOrderSetting.ALL = field.NewField(tableName, "*")
	_omsOrderSetting.ID = field.NewInt64(tableName, "id")
	_omsOrderSetting.FlashOrderOvertime = field.NewInt32(tableName, "flash_order_overtime")
	_omsOrderSetting.NormalOrderOvertime = field.NewInt32(tableName, "normal_order_overtime")
	_omsOrderSetting.ConfirmOvertime = field.NewInt32(tableName, "confirm_overtime")
	_omsOrderSetting.FinishOvertime = field.NewInt32(tableName, "finish_overtime")
	_omsOrderSetting.CommentOvertime = field.NewInt32(tableName, "comment_overtime")

	_omsOrderSetting.fillFieldMap()

	return _omsOrderSetting
}

type omsOrderSetting struct {
	omsOrderSettingDo

	ALL                 field.Field
	ID                  field.Int64
	FlashOrderOvertime  field.Int32
	NormalOrderOvertime field.Int32
	ConfirmOvertime     field.Int32
	FinishOvertime      field.Int32
	CommentOvertime     field.Int32

	fieldMap map[string]field.Expr
}

func (o omsOrderSetting) Table(newTableName string) *omsOrderSetting {
	o.omsOrderSettingDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderSetting) As(alias string) *omsOrderSetting {
	o.omsOrderSettingDo.DO = *(o.omsOrderSettingDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderSetting) updateTableName(table string) *omsOrderSetting {
	o.ALL = field.NewField(table, "*")
	o.ID = field.NewInt64(table, "id")
	o.FlashOrderOvertime = field.NewInt32(table, "flash_order_overtime")
	o.NormalOrderOvertime = field.NewInt32(table, "normal_order_overtime")
	o.ConfirmOvertime = field.NewInt32(table, "confirm_overtime")
	o.FinishOvertime = field.NewInt32(table, "finish_overtime")
	o.CommentOvertime = field.NewInt32(table, "comment_overtime")

	o.fillFieldMap()

	return o
}

func (o *omsOrderSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderSetting) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["flash_order_overtime"] = o.FlashOrderOvertime
	o.fieldMap["normal_order_overtime"] = o.NormalOrderOvertime
	o.fieldMap["confirm_overtime"] = o.ConfirmOvertime
	o.fieldMap["finish_overtime"] = o.FinishOvertime
	o.fieldMap["comment_overtime"] = o.CommentOvertime
}

func (o omsOrderSetting) clone(db *gorm.DB) omsOrderSetting {
	o.omsOrderSettingDo.ReplaceDB(db)
	return o
}

type omsOrderSettingDo struct{ gen.DO }

func (o omsOrderSettingDo) Debug() *omsOrderSettingDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderSettingDo) WithContext(ctx context.Context) *omsOrderSettingDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderSettingDo) ReadDB() *omsOrderSettingDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderSettingDo) WriteDB() *omsOrderSettingDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderSettingDo) Clauses(conds ...clause.Expression) *omsOrderSettingDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderSettingDo) Returning(value interface{}, columns ...string) *omsOrderSettingDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderSettingDo) Not(conds ...gen.Condition) *omsOrderSettingDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderSettingDo) Or(conds ...gen.Condition) *omsOrderSettingDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderSettingDo) Select(conds ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderSettingDo) Where(conds ...gen.Condition) *omsOrderSettingDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderSettingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *omsOrderSettingDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o omsOrderSettingDo) Order(conds ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderSettingDo) Distinct(cols ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderSettingDo) Omit(cols ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderSettingDo) Join(table schema.Tabler, on ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderSettingDo) Group(cols ...field.Expr) *omsOrderSettingDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderSettingDo) Having(conds ...gen.Condition) *omsOrderSettingDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderSettingDo) Limit(limit int) *omsOrderSettingDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderSettingDo) Offset(offset int) *omsOrderSettingDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *omsOrderSettingDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderSettingDo) Unscoped() *omsOrderSettingDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderSettingDo) Create(values ...*model.OmsOrderSetting) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderSettingDo) CreateInBatches(values []*model.OmsOrderSetting, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderSettingDo) Save(values ...*model.OmsOrderSetting) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderSettingDo) First() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) Take() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) Last() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) Find() ([]*model.OmsOrderSetting, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderSetting), err
}

func (o omsOrderSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderSetting, err error) {
	buf := make([]*model.OmsOrderSetting, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderSettingDo) FindInBatches(result *[]*model.OmsOrderSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderSettingDo) Attrs(attrs ...field.AssignExpr) *omsOrderSettingDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderSettingDo) Assign(attrs ...field.AssignExpr) *omsOrderSettingDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderSettingDo) Joins(fields ...field.RelationField) *omsOrderSettingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderSettingDo) Preload(fields ...field.RelationField) *omsOrderSettingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderSettingDo) FirstOrInit() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) FirstOrCreate() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) FindByPage(offset int, limit int) (result []*model.OmsOrderSetting, count int64, err error) {
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

func (o omsOrderSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderSettingDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderSettingDo) Delete(models ...*model.OmsOrderSetting) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderSettingDo) withDO(do gen.Dao) *omsOrderSettingDo {
	o.DO = *do.(*gen.DO)
	return o
}