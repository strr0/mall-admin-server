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

func newSmsCouponHistory(db *gorm.DB) smsCouponHistory {
	_smsCouponHistory := smsCouponHistory{}

	_smsCouponHistory.smsCouponHistoryDo.UseDB(db)
	_smsCouponHistory.smsCouponHistoryDo.UseModel(&model.SmsCouponHistory{})

	tableName := _smsCouponHistory.smsCouponHistoryDo.TableName()
	_smsCouponHistory.ALL = field.NewField(tableName, "*")
	_smsCouponHistory.ID = field.NewInt64(tableName, "id")
	_smsCouponHistory.CouponID = field.NewInt64(tableName, "coupon_id")
	_smsCouponHistory.MemberID = field.NewInt64(tableName, "member_id")
	_smsCouponHistory.CouponCode = field.NewString(tableName, "coupon_code")
	_smsCouponHistory.MemberNickname = field.NewString(tableName, "member_nickname")
	_smsCouponHistory.GetType = field.NewInt32(tableName, "get_type")
	_smsCouponHistory.CreateTime = field.NewTime(tableName, "create_time")
	_smsCouponHistory.UseStatus = field.NewInt32(tableName, "use_status")
	_smsCouponHistory.UseTime = field.NewTime(tableName, "use_time")
	_smsCouponHistory.OrderID = field.NewInt64(tableName, "order_id")
	_smsCouponHistory.OrderSn = field.NewString(tableName, "order_sn")

	_smsCouponHistory.fillFieldMap()

	return _smsCouponHistory
}

type smsCouponHistory struct {
	smsCouponHistoryDo

	ALL            field.Field
	ID             field.Int64
	CouponID       field.Int64
	MemberID       field.Int64
	CouponCode     field.String
	MemberNickname field.String
	GetType        field.Int32
	CreateTime     field.Time
	UseStatus      field.Int32
	UseTime        field.Time
	OrderID        field.Int64
	OrderSn        field.String

	fieldMap map[string]field.Expr
}

func (s smsCouponHistory) Table(newTableName string) *smsCouponHistory {
	s.smsCouponHistoryDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsCouponHistory) As(alias string) *smsCouponHistory {
	s.smsCouponHistoryDo.DO = *(s.smsCouponHistoryDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsCouponHistory) updateTableName(table string) *smsCouponHistory {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.CouponID = field.NewInt64(table, "coupon_id")
	s.MemberID = field.NewInt64(table, "member_id")
	s.CouponCode = field.NewString(table, "coupon_code")
	s.MemberNickname = field.NewString(table, "member_nickname")
	s.GetType = field.NewInt32(table, "get_type")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UseStatus = field.NewInt32(table, "use_status")
	s.UseTime = field.NewTime(table, "use_time")
	s.OrderID = field.NewInt64(table, "order_id")
	s.OrderSn = field.NewString(table, "order_sn")

	s.fillFieldMap()

	return s
}

func (s *smsCouponHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsCouponHistory) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["coupon_id"] = s.CouponID
	s.fieldMap["member_id"] = s.MemberID
	s.fieldMap["coupon_code"] = s.CouponCode
	s.fieldMap["member_nickname"] = s.MemberNickname
	s.fieldMap["get_type"] = s.GetType
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["use_status"] = s.UseStatus
	s.fieldMap["use_time"] = s.UseTime
	s.fieldMap["order_id"] = s.OrderID
	s.fieldMap["order_sn"] = s.OrderSn
}

func (s smsCouponHistory) clone(db *gorm.DB) smsCouponHistory {
	s.smsCouponHistoryDo.ReplaceDB(db)
	return s
}

type smsCouponHistoryDo struct{ gen.DO }

func (s smsCouponHistoryDo) Debug() *smsCouponHistoryDo {
	return s.withDO(s.DO.Debug())
}

func (s smsCouponHistoryDo) WithContext(ctx context.Context) *smsCouponHistoryDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsCouponHistoryDo) ReadDB() *smsCouponHistoryDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsCouponHistoryDo) WriteDB() *smsCouponHistoryDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsCouponHistoryDo) Clauses(conds ...clause.Expression) *smsCouponHistoryDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsCouponHistoryDo) Returning(value interface{}, columns ...string) *smsCouponHistoryDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsCouponHistoryDo) Not(conds ...gen.Condition) *smsCouponHistoryDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsCouponHistoryDo) Or(conds ...gen.Condition) *smsCouponHistoryDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsCouponHistoryDo) Select(conds ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsCouponHistoryDo) Where(conds ...gen.Condition) *smsCouponHistoryDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsCouponHistoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *smsCouponHistoryDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s smsCouponHistoryDo) Order(conds ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsCouponHistoryDo) Distinct(cols ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsCouponHistoryDo) Omit(cols ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsCouponHistoryDo) Join(table schema.Tabler, on ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsCouponHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsCouponHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsCouponHistoryDo) Group(cols ...field.Expr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsCouponHistoryDo) Having(conds ...gen.Condition) *smsCouponHistoryDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsCouponHistoryDo) Limit(limit int) *smsCouponHistoryDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsCouponHistoryDo) Offset(offset int) *smsCouponHistoryDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsCouponHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *smsCouponHistoryDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsCouponHistoryDo) Unscoped() *smsCouponHistoryDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsCouponHistoryDo) Create(values ...*model.SmsCouponHistory) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsCouponHistoryDo) CreateInBatches(values []*model.SmsCouponHistory, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsCouponHistoryDo) Save(values ...*model.SmsCouponHistory) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsCouponHistoryDo) First() (*model.SmsCouponHistory, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponHistory), nil
	}
}

func (s smsCouponHistoryDo) Take() (*model.SmsCouponHistory, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponHistory), nil
	}
}

func (s smsCouponHistoryDo) Last() (*model.SmsCouponHistory, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponHistory), nil
	}
}

func (s smsCouponHistoryDo) Find() ([]*model.SmsCouponHistory, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsCouponHistory), err
}

func (s smsCouponHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsCouponHistory, err error) {
	buf := make([]*model.SmsCouponHistory, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsCouponHistoryDo) FindInBatches(result *[]*model.SmsCouponHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsCouponHistoryDo) Attrs(attrs ...field.AssignExpr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsCouponHistoryDo) Assign(attrs ...field.AssignExpr) *smsCouponHistoryDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsCouponHistoryDo) Joins(fields ...field.RelationField) *smsCouponHistoryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsCouponHistoryDo) Preload(fields ...field.RelationField) *smsCouponHistoryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsCouponHistoryDo) FirstOrInit() (*model.SmsCouponHistory, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponHistory), nil
	}
}

func (s smsCouponHistoryDo) FirstOrCreate() (*model.SmsCouponHistory, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponHistory), nil
	}
}

func (s smsCouponHistoryDo) FindByPage(offset int, limit int) (result []*model.SmsCouponHistory, count int64, err error) {
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

func (s smsCouponHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsCouponHistoryDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsCouponHistoryDo) Delete(models ...*model.SmsCouponHistory) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsCouponHistoryDo) withDO(do gen.Dao) *smsCouponHistoryDo {
	s.DO = *do.(*gen.DO)
	return s
}
