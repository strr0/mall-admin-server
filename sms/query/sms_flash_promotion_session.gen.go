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

func newSmsFlashPromotionSession(db *gorm.DB) smsFlashPromotionSession {
	_smsFlashPromotionSession := smsFlashPromotionSession{}

	_smsFlashPromotionSession.smsFlashPromotionSessionDo.UseDB(db)
	_smsFlashPromotionSession.smsFlashPromotionSessionDo.UseModel(&model.SmsFlashPromotionSession{})

	tableName := _smsFlashPromotionSession.smsFlashPromotionSessionDo.TableName()
	_smsFlashPromotionSession.ALL = field.NewField(tableName, "*")
	_smsFlashPromotionSession.ID = field.NewInt64(tableName, "id")
	_smsFlashPromotionSession.Name = field.NewString(tableName, "name")
	_smsFlashPromotionSession.StartTime = field.NewTime(tableName, "start_time")
	_smsFlashPromotionSession.EndTime = field.NewTime(tableName, "end_time")
	_smsFlashPromotionSession.Status = field.NewInt32(tableName, "status")
	_smsFlashPromotionSession.CreateTime = field.NewTime(tableName, "create_time")

	_smsFlashPromotionSession.fillFieldMap()

	return _smsFlashPromotionSession
}

type smsFlashPromotionSession struct {
	smsFlashPromotionSessionDo

	ALL        field.Field
	ID         field.Int64
	Name       field.String
	StartTime  field.Time
	EndTime    field.Time
	Status     field.Int32
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (s smsFlashPromotionSession) Table(newTableName string) *smsFlashPromotionSession {
	s.smsFlashPromotionSessionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsFlashPromotionSession) As(alias string) *smsFlashPromotionSession {
	s.smsFlashPromotionSessionDo.DO = *(s.smsFlashPromotionSessionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsFlashPromotionSession) updateTableName(table string) *smsFlashPromotionSession {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.StartTime = field.NewTime(table, "start_time")
	s.EndTime = field.NewTime(table, "end_time")
	s.Status = field.NewInt32(table, "status")
	s.CreateTime = field.NewTime(table, "create_time")

	s.fillFieldMap()

	return s
}

func (s *smsFlashPromotionSession) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsFlashPromotionSession) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["start_time"] = s.StartTime
	s.fieldMap["end_time"] = s.EndTime
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_time"] = s.CreateTime
}

func (s smsFlashPromotionSession) clone(db *gorm.DB) smsFlashPromotionSession {
	s.smsFlashPromotionSessionDo.ReplaceDB(db)
	return s
}

type smsFlashPromotionSessionDo struct{ gen.DO }

func (s smsFlashPromotionSessionDo) Debug() *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Debug())
}

func (s smsFlashPromotionSessionDo) WithContext(ctx context.Context) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsFlashPromotionSessionDo) ReadDB() *smsFlashPromotionSessionDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsFlashPromotionSessionDo) WriteDB() *smsFlashPromotionSessionDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsFlashPromotionSessionDo) Clauses(conds ...clause.Expression) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsFlashPromotionSessionDo) Returning(value interface{}, columns ...string) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsFlashPromotionSessionDo) Not(conds ...gen.Condition) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsFlashPromotionSessionDo) Or(conds ...gen.Condition) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsFlashPromotionSessionDo) Select(conds ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsFlashPromotionSessionDo) Where(conds ...gen.Condition) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsFlashPromotionSessionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *smsFlashPromotionSessionDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s smsFlashPromotionSessionDo) Order(conds ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsFlashPromotionSessionDo) Distinct(cols ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsFlashPromotionSessionDo) Omit(cols ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsFlashPromotionSessionDo) Join(table schema.Tabler, on ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsFlashPromotionSessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsFlashPromotionSessionDo) RightJoin(table schema.Tabler, on ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsFlashPromotionSessionDo) Group(cols ...field.Expr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsFlashPromotionSessionDo) Having(conds ...gen.Condition) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsFlashPromotionSessionDo) Limit(limit int) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsFlashPromotionSessionDo) Offset(offset int) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsFlashPromotionSessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsFlashPromotionSessionDo) Unscoped() *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsFlashPromotionSessionDo) Create(values ...*model.SmsFlashPromotionSession) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsFlashPromotionSessionDo) CreateInBatches(values []*model.SmsFlashPromotionSession, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsFlashPromotionSessionDo) Save(values ...*model.SmsFlashPromotionSession) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsFlashPromotionSessionDo) First() (*model.SmsFlashPromotionSession, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotionSession), nil
	}
}

func (s smsFlashPromotionSessionDo) Take() (*model.SmsFlashPromotionSession, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotionSession), nil
	}
}

func (s smsFlashPromotionSessionDo) Last() (*model.SmsFlashPromotionSession, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotionSession), nil
	}
}

func (s smsFlashPromotionSessionDo) Find() ([]*model.SmsFlashPromotionSession, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsFlashPromotionSession), err
}

func (s smsFlashPromotionSessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsFlashPromotionSession, err error) {
	buf := make([]*model.SmsFlashPromotionSession, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsFlashPromotionSessionDo) FindInBatches(result *[]*model.SmsFlashPromotionSession, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsFlashPromotionSessionDo) Attrs(attrs ...field.AssignExpr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsFlashPromotionSessionDo) Assign(attrs ...field.AssignExpr) *smsFlashPromotionSessionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsFlashPromotionSessionDo) Joins(fields ...field.RelationField) *smsFlashPromotionSessionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsFlashPromotionSessionDo) Preload(fields ...field.RelationField) *smsFlashPromotionSessionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsFlashPromotionSessionDo) FirstOrInit() (*model.SmsFlashPromotionSession, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotionSession), nil
	}
}

func (s smsFlashPromotionSessionDo) FirstOrCreate() (*model.SmsFlashPromotionSession, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsFlashPromotionSession), nil
	}
}

func (s smsFlashPromotionSessionDo) FindByPage(offset int, limit int) (result []*model.SmsFlashPromotionSession, count int64, err error) {
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

func (s smsFlashPromotionSessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsFlashPromotionSessionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsFlashPromotionSessionDo) Delete(models ...*model.SmsFlashPromotionSession) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsFlashPromotionSessionDo) withDO(do gen.Dao) *smsFlashPromotionSessionDo {
	s.DO = *do.(*gen.DO)
	return s
}