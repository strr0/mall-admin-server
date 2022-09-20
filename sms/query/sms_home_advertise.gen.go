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

func newSmsHomeAdvertise(db *gorm.DB) smsHomeAdvertise {
	_smsHomeAdvertise := smsHomeAdvertise{}

	_smsHomeAdvertise.smsHomeAdvertiseDo.UseDB(db)
	_smsHomeAdvertise.smsHomeAdvertiseDo.UseModel(&model.SmsHomeAdvertise{})

	tableName := _smsHomeAdvertise.smsHomeAdvertiseDo.TableName()
	_smsHomeAdvertise.ALL = field.NewField(tableName, "*")
	_smsHomeAdvertise.ID = field.NewInt64(tableName, "id")
	_smsHomeAdvertise.Name = field.NewString(tableName, "name")
	_smsHomeAdvertise.Type = field.NewInt32(tableName, "type")
	_smsHomeAdvertise.Pic = field.NewString(tableName, "pic")
	_smsHomeAdvertise.StartTime = field.NewTime(tableName, "start_time")
	_smsHomeAdvertise.EndTime = field.NewTime(tableName, "end_time")
	_smsHomeAdvertise.Status = field.NewInt32(tableName, "status")
	_smsHomeAdvertise.ClickCount = field.NewInt32(tableName, "click_count")
	_smsHomeAdvertise.OrderCount = field.NewInt32(tableName, "order_count")
	_smsHomeAdvertise.URL = field.NewString(tableName, "url")
	_smsHomeAdvertise.Note = field.NewString(tableName, "note")
	_smsHomeAdvertise.Sort = field.NewInt32(tableName, "sort")

	_smsHomeAdvertise.fillFieldMap()

	return _smsHomeAdvertise
}

type smsHomeAdvertise struct {
	smsHomeAdvertiseDo

	ALL        field.Field
	ID         field.Int64
	Name       field.String
	Type       field.Int32
	Pic        field.String
	StartTime  field.Time
	EndTime    field.Time
	Status     field.Int32
	ClickCount field.Int32
	OrderCount field.Int32
	URL        field.String
	Note       field.String
	Sort       field.Int32

	fieldMap map[string]field.Expr
}

func (s smsHomeAdvertise) Table(newTableName string) *smsHomeAdvertise {
	s.smsHomeAdvertiseDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsHomeAdvertise) As(alias string) *smsHomeAdvertise {
	s.smsHomeAdvertiseDo.DO = *(s.smsHomeAdvertiseDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsHomeAdvertise) updateTableName(table string) *smsHomeAdvertise {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Type = field.NewInt32(table, "type")
	s.Pic = field.NewString(table, "pic")
	s.StartTime = field.NewTime(table, "start_time")
	s.EndTime = field.NewTime(table, "end_time")
	s.Status = field.NewInt32(table, "status")
	s.ClickCount = field.NewInt32(table, "click_count")
	s.OrderCount = field.NewInt32(table, "order_count")
	s.URL = field.NewString(table, "url")
	s.Note = field.NewString(table, "note")
	s.Sort = field.NewInt32(table, "sort")

	s.fillFieldMap()

	return s
}

func (s *smsHomeAdvertise) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsHomeAdvertise) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["type"] = s.Type
	s.fieldMap["pic"] = s.Pic
	s.fieldMap["start_time"] = s.StartTime
	s.fieldMap["end_time"] = s.EndTime
	s.fieldMap["status"] = s.Status
	s.fieldMap["click_count"] = s.ClickCount
	s.fieldMap["order_count"] = s.OrderCount
	s.fieldMap["url"] = s.URL
	s.fieldMap["note"] = s.Note
	s.fieldMap["sort"] = s.Sort
}

func (s smsHomeAdvertise) clone(db *gorm.DB) smsHomeAdvertise {
	s.smsHomeAdvertiseDo.ReplaceDB(db)
	return s
}

type smsHomeAdvertiseDo struct{ gen.DO }

func (s smsHomeAdvertiseDo) Debug() *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Debug())
}

func (s smsHomeAdvertiseDo) WithContext(ctx context.Context) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsHomeAdvertiseDo) ReadDB() *smsHomeAdvertiseDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsHomeAdvertiseDo) WriteDB() *smsHomeAdvertiseDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsHomeAdvertiseDo) Clauses(conds ...clause.Expression) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsHomeAdvertiseDo) Returning(value interface{}, columns ...string) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsHomeAdvertiseDo) Not(conds ...gen.Condition) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsHomeAdvertiseDo) Or(conds ...gen.Condition) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsHomeAdvertiseDo) Select(conds ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsHomeAdvertiseDo) Where(conds ...gen.Condition) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsHomeAdvertiseDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *smsHomeAdvertiseDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s smsHomeAdvertiseDo) Order(conds ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsHomeAdvertiseDo) Distinct(cols ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsHomeAdvertiseDo) Omit(cols ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsHomeAdvertiseDo) Join(table schema.Tabler, on ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsHomeAdvertiseDo) LeftJoin(table schema.Tabler, on ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsHomeAdvertiseDo) RightJoin(table schema.Tabler, on ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsHomeAdvertiseDo) Group(cols ...field.Expr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsHomeAdvertiseDo) Having(conds ...gen.Condition) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsHomeAdvertiseDo) Limit(limit int) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsHomeAdvertiseDo) Offset(offset int) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsHomeAdvertiseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsHomeAdvertiseDo) Unscoped() *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsHomeAdvertiseDo) Create(values ...*model.SmsHomeAdvertise) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsHomeAdvertiseDo) CreateInBatches(values []*model.SmsHomeAdvertise, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsHomeAdvertiseDo) Save(values ...*model.SmsHomeAdvertise) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsHomeAdvertiseDo) First() (*model.SmsHomeAdvertise, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdvertise), nil
	}
}

func (s smsHomeAdvertiseDo) Take() (*model.SmsHomeAdvertise, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdvertise), nil
	}
}

func (s smsHomeAdvertiseDo) Last() (*model.SmsHomeAdvertise, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdvertise), nil
	}
}

func (s smsHomeAdvertiseDo) Find() ([]*model.SmsHomeAdvertise, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsHomeAdvertise), err
}

func (s smsHomeAdvertiseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsHomeAdvertise, err error) {
	buf := make([]*model.SmsHomeAdvertise, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsHomeAdvertiseDo) FindInBatches(result *[]*model.SmsHomeAdvertise, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsHomeAdvertiseDo) Attrs(attrs ...field.AssignExpr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsHomeAdvertiseDo) Assign(attrs ...field.AssignExpr) *smsHomeAdvertiseDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsHomeAdvertiseDo) Joins(fields ...field.RelationField) *smsHomeAdvertiseDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsHomeAdvertiseDo) Preload(fields ...field.RelationField) *smsHomeAdvertiseDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsHomeAdvertiseDo) FirstOrInit() (*model.SmsHomeAdvertise, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdvertise), nil
	}
}

func (s smsHomeAdvertiseDo) FirstOrCreate() (*model.SmsHomeAdvertise, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdvertise), nil
	}
}

func (s smsHomeAdvertiseDo) FindByPage(offset int, limit int) (result []*model.SmsHomeAdvertise, count int64, err error) {
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

func (s smsHomeAdvertiseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsHomeAdvertiseDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsHomeAdvertiseDo) Delete(models ...*model.SmsHomeAdvertise) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsHomeAdvertiseDo) withDO(do gen.Dao) *smsHomeAdvertiseDo {
	s.DO = *do.(*gen.DO)
	return s
}
