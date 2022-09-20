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

	"mall-admin-server/sms/model"
)

func newSmsCouponProductRelation(db *gorm.DB) smsCouponProductRelation {
	_smsCouponProductRelation := smsCouponProductRelation{}

	_smsCouponProductRelation.smsCouponProductRelationDo.UseDB(db)
	_smsCouponProductRelation.smsCouponProductRelationDo.UseModel(&model.SmsCouponProductRelation{})

	tableName := _smsCouponProductRelation.smsCouponProductRelationDo.TableName()
	_smsCouponProductRelation.ALL = field.NewField(tableName, "*")
	_smsCouponProductRelation.ID = field.NewInt64(tableName, "id")
	_smsCouponProductRelation.CouponID = field.NewInt64(tableName, "coupon_id")
	_smsCouponProductRelation.ProductID = field.NewInt64(tableName, "product_id")
	_smsCouponProductRelation.ProductName = field.NewString(tableName, "product_name")
	_smsCouponProductRelation.ProductSn = field.NewString(tableName, "product_sn")

	_smsCouponProductRelation.fillFieldMap()

	return _smsCouponProductRelation
}

type smsCouponProductRelation struct {
	smsCouponProductRelationDo

	ALL         field.Field
	ID          field.Int64
	CouponID    field.Int64
	ProductID   field.Int64
	ProductName field.String
	ProductSn   field.String

	fieldMap map[string]field.Expr
}

func (s smsCouponProductRelation) Table(newTableName string) *smsCouponProductRelation {
	s.smsCouponProductRelationDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsCouponProductRelation) As(alias string) *smsCouponProductRelation {
	s.smsCouponProductRelationDo.DO = *(s.smsCouponProductRelationDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsCouponProductRelation) updateTableName(table string) *smsCouponProductRelation {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.CouponID = field.NewInt64(table, "coupon_id")
	s.ProductID = field.NewInt64(table, "product_id")
	s.ProductName = field.NewString(table, "product_name")
	s.ProductSn = field.NewString(table, "product_sn")

	s.fillFieldMap()

	return s
}

func (s *smsCouponProductRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsCouponProductRelation) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id"] = s.ID
	s.fieldMap["coupon_id"] = s.CouponID
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["product_name"] = s.ProductName
	s.fieldMap["product_sn"] = s.ProductSn
}

func (s smsCouponProductRelation) clone(db *gorm.DB) smsCouponProductRelation {
	s.smsCouponProductRelationDo.ReplaceDB(db)
	return s
}

type smsCouponProductRelationDo struct{ gen.DO }

func (s smsCouponProductRelationDo) Debug() *smsCouponProductRelationDo {
	return s.withDO(s.DO.Debug())
}

func (s smsCouponProductRelationDo) WithContext(ctx context.Context) *smsCouponProductRelationDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsCouponProductRelationDo) ReadDB() *smsCouponProductRelationDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsCouponProductRelationDo) WriteDB() *smsCouponProductRelationDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsCouponProductRelationDo) Clauses(conds ...clause.Expression) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsCouponProductRelationDo) Returning(value interface{}, columns ...string) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsCouponProductRelationDo) Not(conds ...gen.Condition) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsCouponProductRelationDo) Or(conds ...gen.Condition) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsCouponProductRelationDo) Select(conds ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsCouponProductRelationDo) Where(conds ...gen.Condition) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsCouponProductRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *smsCouponProductRelationDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s smsCouponProductRelationDo) Order(conds ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsCouponProductRelationDo) Distinct(cols ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsCouponProductRelationDo) Omit(cols ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsCouponProductRelationDo) Join(table schema.Tabler, on ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsCouponProductRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsCouponProductRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsCouponProductRelationDo) Group(cols ...field.Expr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsCouponProductRelationDo) Having(conds ...gen.Condition) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsCouponProductRelationDo) Limit(limit int) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsCouponProductRelationDo) Offset(offset int) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsCouponProductRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsCouponProductRelationDo) Unscoped() *smsCouponProductRelationDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsCouponProductRelationDo) Create(values ...*model.SmsCouponProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsCouponProductRelationDo) CreateInBatches(values []*model.SmsCouponProductRelation, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsCouponProductRelationDo) Save(values ...*model.SmsCouponProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsCouponProductRelationDo) First() (*model.SmsCouponProductRelation, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponProductRelation), nil
	}
}

func (s smsCouponProductRelationDo) Take() (*model.SmsCouponProductRelation, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponProductRelation), nil
	}
}

func (s smsCouponProductRelationDo) Last() (*model.SmsCouponProductRelation, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponProductRelation), nil
	}
}

func (s smsCouponProductRelationDo) Find() ([]*model.SmsCouponProductRelation, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsCouponProductRelation), err
}

func (s smsCouponProductRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsCouponProductRelation, err error) {
	buf := make([]*model.SmsCouponProductRelation, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsCouponProductRelationDo) FindInBatches(result *[]*model.SmsCouponProductRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsCouponProductRelationDo) Attrs(attrs ...field.AssignExpr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsCouponProductRelationDo) Assign(attrs ...field.AssignExpr) *smsCouponProductRelationDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsCouponProductRelationDo) Joins(fields ...field.RelationField) *smsCouponProductRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsCouponProductRelationDo) Preload(fields ...field.RelationField) *smsCouponProductRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsCouponProductRelationDo) FirstOrInit() (*model.SmsCouponProductRelation, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponProductRelation), nil
	}
}

func (s smsCouponProductRelationDo) FirstOrCreate() (*model.SmsCouponProductRelation, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponProductRelation), nil
	}
}

func (s smsCouponProductRelationDo) FindByPage(offset int, limit int) (result []*model.SmsCouponProductRelation, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s smsCouponProductRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsCouponProductRelationDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsCouponProductRelationDo) Delete(models ...*model.SmsCouponProductRelation) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsCouponProductRelationDo) withDO(do gen.Dao) *smsCouponProductRelationDo {
	s.DO = *do.(*gen.DO)
	return s
}