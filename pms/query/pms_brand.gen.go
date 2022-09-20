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

func newPmsBrand(db *gorm.DB) pmsBrand {
	_pmsBrand := pmsBrand{}

	_pmsBrand.pmsBrandDo.UseDB(db)
	_pmsBrand.pmsBrandDo.UseModel(&model.PmsBrand{})

	tableName := _pmsBrand.pmsBrandDo.TableName()
	_pmsBrand.ALL = field.NewField(tableName, "*")
	_pmsBrand.ID = field.NewInt64(tableName, "id")
	_pmsBrand.Name = field.NewString(tableName, "name")
	_pmsBrand.FirstLetter = field.NewString(tableName, "first_letter")
	_pmsBrand.Sort = field.NewInt32(tableName, "sort")
	_pmsBrand.FactoryStatus = field.NewInt32(tableName, "factory_status")
	_pmsBrand.ShowStatus = field.NewInt32(tableName, "show_status")
	_pmsBrand.ProductCount = field.NewInt32(tableName, "product_count")
	_pmsBrand.ProductCommentCount = field.NewInt32(tableName, "product_comment_count")
	_pmsBrand.Logo = field.NewString(tableName, "logo")
	_pmsBrand.BigPic = field.NewString(tableName, "big_pic")
	_pmsBrand.BrandStory = field.NewString(tableName, "brand_story")

	_pmsBrand.fillFieldMap()

	return _pmsBrand
}

type pmsBrand struct {
	pmsBrandDo

	ALL                 field.Field
	ID                  field.Int64
	Name                field.String
	FirstLetter         field.String
	Sort                field.Int32
	FactoryStatus       field.Int32
	ShowStatus          field.Int32
	ProductCount        field.Int32
	ProductCommentCount field.Int32
	Logo                field.String
	BigPic              field.String
	BrandStory          field.String

	fieldMap map[string]field.Expr
}

func (p pmsBrand) Table(newTableName string) *pmsBrand {
	p.pmsBrandDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsBrand) As(alias string) *pmsBrand {
	p.pmsBrandDo.DO = *(p.pmsBrandDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsBrand) updateTableName(table string) *pmsBrand {
	p.ALL = field.NewField(table, "*")
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.FirstLetter = field.NewString(table, "first_letter")
	p.Sort = field.NewInt32(table, "sort")
	p.FactoryStatus = field.NewInt32(table, "factory_status")
	p.ShowStatus = field.NewInt32(table, "show_status")
	p.ProductCount = field.NewInt32(table, "product_count")
	p.ProductCommentCount = field.NewInt32(table, "product_comment_count")
	p.Logo = field.NewString(table, "logo")
	p.BigPic = field.NewString(table, "big_pic")
	p.BrandStory = field.NewString(table, "brand_story")

	p.fillFieldMap()

	return p
}

func (p *pmsBrand) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsBrand) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["first_letter"] = p.FirstLetter
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["factory_status"] = p.FactoryStatus
	p.fieldMap["show_status"] = p.ShowStatus
	p.fieldMap["product_count"] = p.ProductCount
	p.fieldMap["product_comment_count"] = p.ProductCommentCount
	p.fieldMap["logo"] = p.Logo
	p.fieldMap["big_pic"] = p.BigPic
	p.fieldMap["brand_story"] = p.BrandStory
}

func (p pmsBrand) clone(db *gorm.DB) pmsBrand {
	p.pmsBrandDo.ReplaceDB(db)
	return p
}

type pmsBrandDo struct{ gen.DO }

func (p pmsBrandDo) Debug() *pmsBrandDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsBrandDo) WithContext(ctx context.Context) *pmsBrandDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsBrandDo) ReadDB() *pmsBrandDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsBrandDo) WriteDB() *pmsBrandDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsBrandDo) Clauses(conds ...clause.Expression) *pmsBrandDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsBrandDo) Returning(value interface{}, columns ...string) *pmsBrandDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsBrandDo) Not(conds ...gen.Condition) *pmsBrandDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsBrandDo) Or(conds ...gen.Condition) *pmsBrandDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsBrandDo) Select(conds ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsBrandDo) Where(conds ...gen.Condition) *pmsBrandDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsBrandDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pmsBrandDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pmsBrandDo) Order(conds ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsBrandDo) Distinct(cols ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsBrandDo) Omit(cols ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsBrandDo) Join(table schema.Tabler, on ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsBrandDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsBrandDo) RightJoin(table schema.Tabler, on ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsBrandDo) Group(cols ...field.Expr) *pmsBrandDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsBrandDo) Having(conds ...gen.Condition) *pmsBrandDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsBrandDo) Limit(limit int) *pmsBrandDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsBrandDo) Offset(offset int) *pmsBrandDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsBrandDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pmsBrandDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsBrandDo) Unscoped() *pmsBrandDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsBrandDo) Create(values ...*model.PmsBrand) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsBrandDo) CreateInBatches(values []*model.PmsBrand, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsBrandDo) Save(values ...*model.PmsBrand) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsBrandDo) First() (*model.PmsBrand, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsBrand), nil
	}
}

func (p pmsBrandDo) Take() (*model.PmsBrand, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsBrand), nil
	}
}

func (p pmsBrandDo) Last() (*model.PmsBrand, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsBrand), nil
	}
}

func (p pmsBrandDo) Find() ([]*model.PmsBrand, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsBrand), err
}

func (p pmsBrandDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsBrand, err error) {
	buf := make([]*model.PmsBrand, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsBrandDo) FindInBatches(result *[]*model.PmsBrand, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsBrandDo) Attrs(attrs ...field.AssignExpr) *pmsBrandDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsBrandDo) Assign(attrs ...field.AssignExpr) *pmsBrandDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsBrandDo) Joins(fields ...field.RelationField) *pmsBrandDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsBrandDo) Preload(fields ...field.RelationField) *pmsBrandDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsBrandDo) FirstOrInit() (*model.PmsBrand, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsBrand), nil
	}
}

func (p pmsBrandDo) FirstOrCreate() (*model.PmsBrand, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsBrand), nil
	}
}

func (p pmsBrandDo) FindByPage(offset int, limit int) (result []*model.PmsBrand, count int64, err error) {
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

func (p pmsBrandDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsBrandDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsBrandDo) Delete(models ...*model.PmsBrand) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsBrandDo) withDO(do gen.Dao) *pmsBrandDo {
	p.DO = *do.(*gen.DO)
	return p
}
