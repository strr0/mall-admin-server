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

	"mall-admin-server/ums/model"
)

func newUmsResourceCategory(db *gorm.DB) umsResourceCategory {
	_umsResourceCategory := umsResourceCategory{}

	_umsResourceCategory.umsResourceCategoryDo.UseDB(db)
	_umsResourceCategory.umsResourceCategoryDo.UseModel(&model.UmsResourceCategory{})

	tableName := _umsResourceCategory.umsResourceCategoryDo.TableName()
	_umsResourceCategory.ALL = field.NewField(tableName, "*")
	_umsResourceCategory.ID = field.NewInt64(tableName, "id")
	_umsResourceCategory.CreateTime = field.NewTime(tableName, "create_time")
	_umsResourceCategory.Name = field.NewString(tableName, "name")
	_umsResourceCategory.Sort = field.NewInt32(tableName, "sort")

	_umsResourceCategory.fillFieldMap()

	return _umsResourceCategory
}

type umsResourceCategory struct {
	umsResourceCategoryDo

	ALL        field.Field
	ID         field.Int64
	CreateTime field.Time
	Name       field.String
	Sort       field.Int32

	fieldMap map[string]field.Expr
}

func (u umsResourceCategory) Table(newTableName string) *umsResourceCategory {
	u.umsResourceCategoryDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsResourceCategory) As(alias string) *umsResourceCategory {
	u.umsResourceCategoryDo.DO = *(u.umsResourceCategoryDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsResourceCategory) updateTableName(table string) *umsResourceCategory {
	u.ALL = field.NewField(table, "*")
	u.ID = field.NewInt64(table, "id")
	u.CreateTime = field.NewTime(table, "create_time")
	u.Name = field.NewString(table, "name")
	u.Sort = field.NewInt32(table, "sort")

	u.fillFieldMap()

	return u
}

func (u *umsResourceCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsResourceCategory) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["id"] = u.ID
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["name"] = u.Name
	u.fieldMap["sort"] = u.Sort
}

func (u umsResourceCategory) clone(db *gorm.DB) umsResourceCategory {
	u.umsResourceCategoryDo.ReplaceDB(db)
	return u
}

type umsResourceCategoryDo struct{ gen.DO }

func (u umsResourceCategoryDo) Debug() *umsResourceCategoryDo {
	return u.withDO(u.DO.Debug())
}

func (u umsResourceCategoryDo) WithContext(ctx context.Context) *umsResourceCategoryDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsResourceCategoryDo) ReadDB() *umsResourceCategoryDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsResourceCategoryDo) WriteDB() *umsResourceCategoryDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsResourceCategoryDo) Clauses(conds ...clause.Expression) *umsResourceCategoryDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsResourceCategoryDo) Returning(value interface{}, columns ...string) *umsResourceCategoryDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsResourceCategoryDo) Not(conds ...gen.Condition) *umsResourceCategoryDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsResourceCategoryDo) Or(conds ...gen.Condition) *umsResourceCategoryDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsResourceCategoryDo) Select(conds ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsResourceCategoryDo) Where(conds ...gen.Condition) *umsResourceCategoryDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsResourceCategoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *umsResourceCategoryDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u umsResourceCategoryDo) Order(conds ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsResourceCategoryDo) Distinct(cols ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsResourceCategoryDo) Omit(cols ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsResourceCategoryDo) Join(table schema.Tabler, on ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsResourceCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsResourceCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsResourceCategoryDo) Group(cols ...field.Expr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsResourceCategoryDo) Having(conds ...gen.Condition) *umsResourceCategoryDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsResourceCategoryDo) Limit(limit int) *umsResourceCategoryDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsResourceCategoryDo) Offset(offset int) *umsResourceCategoryDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsResourceCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *umsResourceCategoryDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsResourceCategoryDo) Unscoped() *umsResourceCategoryDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsResourceCategoryDo) Create(values ...*model.UmsResourceCategory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsResourceCategoryDo) CreateInBatches(values []*model.UmsResourceCategory, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsResourceCategoryDo) Save(values ...*model.UmsResourceCategory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsResourceCategoryDo) First() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) Take() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) Last() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) Find() ([]*model.UmsResourceCategory, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsResourceCategory), err
}

func (u umsResourceCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsResourceCategory, err error) {
	buf := make([]*model.UmsResourceCategory, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsResourceCategoryDo) FindInBatches(result *[]*model.UmsResourceCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsResourceCategoryDo) Attrs(attrs ...field.AssignExpr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsResourceCategoryDo) Assign(attrs ...field.AssignExpr) *umsResourceCategoryDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsResourceCategoryDo) Joins(fields ...field.RelationField) *umsResourceCategoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsResourceCategoryDo) Preload(fields ...field.RelationField) *umsResourceCategoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsResourceCategoryDo) FirstOrInit() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) FirstOrCreate() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) FindByPage(offset int, limit int) (result []*model.UmsResourceCategory, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u umsResourceCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsResourceCategoryDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsResourceCategoryDo) Delete(models ...*model.UmsResourceCategory) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsResourceCategoryDo) withDO(do gen.Dao) *umsResourceCategoryDo {
	u.DO = *do.(*gen.DO)
	return u
}