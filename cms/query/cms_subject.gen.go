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

	"mall-admin-server/cms/model"
)

func newCmsSubject(db *gorm.DB) cmsSubject {
	_cmsSubject := cmsSubject{}

	_cmsSubject.cmsSubjectDo.UseDB(db)
	_cmsSubject.cmsSubjectDo.UseModel(&model.CmsSubject{})

	tableName := _cmsSubject.cmsSubjectDo.TableName()
	_cmsSubject.ALL = field.NewField(tableName, "*")
	_cmsSubject.ID = field.NewInt64(tableName, "id")
	_cmsSubject.CategoryID = field.NewInt64(tableName, "category_id")
	_cmsSubject.Title = field.NewString(tableName, "title")
	_cmsSubject.Pic = field.NewString(tableName, "pic")
	_cmsSubject.ProductCount = field.NewInt32(tableName, "product_count")
	_cmsSubject.RecommendStatus = field.NewInt32(tableName, "recommend_status")
	_cmsSubject.CreateTime = field.NewTime(tableName, "create_time")
	_cmsSubject.CollectCount = field.NewInt32(tableName, "collect_count")
	_cmsSubject.ReadCount = field.NewInt32(tableName, "read_count")
	_cmsSubject.CommentCount = field.NewInt32(tableName, "comment_count")
	_cmsSubject.AlbumPics = field.NewString(tableName, "album_pics")
	_cmsSubject.Description = field.NewString(tableName, "description")
	_cmsSubject.ShowStatus = field.NewInt32(tableName, "show_status")
	_cmsSubject.Content = field.NewString(tableName, "content")
	_cmsSubject.ForwardCount = field.NewInt32(tableName, "forward_count")
	_cmsSubject.CategoryName = field.NewString(tableName, "category_name")

	_cmsSubject.fillFieldMap()

	return _cmsSubject
}

type cmsSubject struct {
	cmsSubjectDo

	ALL             field.Field
	ID              field.Int64
	CategoryID      field.Int64
	Title           field.String
	Pic             field.String
	ProductCount    field.Int32
	RecommendStatus field.Int32
	CreateTime      field.Time
	CollectCount    field.Int32
	ReadCount       field.Int32
	CommentCount    field.Int32
	AlbumPics       field.String
	Description     field.String
	ShowStatus      field.Int32
	Content         field.String
	ForwardCount    field.Int32
	CategoryName    field.String

	fieldMap map[string]field.Expr
}

func (c cmsSubject) Table(newTableName string) *cmsSubject {
	c.cmsSubjectDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cmsSubject) As(alias string) *cmsSubject {
	c.cmsSubjectDo.DO = *(c.cmsSubjectDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cmsSubject) updateTableName(table string) *cmsSubject {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt64(table, "id")
	c.CategoryID = field.NewInt64(table, "category_id")
	c.Title = field.NewString(table, "title")
	c.Pic = field.NewString(table, "pic")
	c.ProductCount = field.NewInt32(table, "product_count")
	c.RecommendStatus = field.NewInt32(table, "recommend_status")
	c.CreateTime = field.NewTime(table, "create_time")
	c.CollectCount = field.NewInt32(table, "collect_count")
	c.ReadCount = field.NewInt32(table, "read_count")
	c.CommentCount = field.NewInt32(table, "comment_count")
	c.AlbumPics = field.NewString(table, "album_pics")
	c.Description = field.NewString(table, "description")
	c.ShowStatus = field.NewInt32(table, "show_status")
	c.Content = field.NewString(table, "content")
	c.ForwardCount = field.NewInt32(table, "forward_count")
	c.CategoryName = field.NewString(table, "category_name")

	c.fillFieldMap()

	return c
}

func (c *cmsSubject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cmsSubject) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 16)
	c.fieldMap["id"] = c.ID
	c.fieldMap["category_id"] = c.CategoryID
	c.fieldMap["title"] = c.Title
	c.fieldMap["pic"] = c.Pic
	c.fieldMap["product_count"] = c.ProductCount
	c.fieldMap["recommend_status"] = c.RecommendStatus
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["collect_count"] = c.CollectCount
	c.fieldMap["read_count"] = c.ReadCount
	c.fieldMap["comment_count"] = c.CommentCount
	c.fieldMap["album_pics"] = c.AlbumPics
	c.fieldMap["description"] = c.Description
	c.fieldMap["show_status"] = c.ShowStatus
	c.fieldMap["content"] = c.Content
	c.fieldMap["forward_count"] = c.ForwardCount
	c.fieldMap["category_name"] = c.CategoryName
}

func (c cmsSubject) clone(db *gorm.DB) cmsSubject {
	c.cmsSubjectDo.ReplaceDB(db)
	return c
}

type cmsSubjectDo struct{ gen.DO }

func (c cmsSubjectDo) Debug() *cmsSubjectDo {
	return c.withDO(c.DO.Debug())
}

func (c cmsSubjectDo) WithContext(ctx context.Context) *cmsSubjectDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cmsSubjectDo) ReadDB() *cmsSubjectDo {
	return c.Clauses(dbresolver.Read)
}

func (c cmsSubjectDo) WriteDB() *cmsSubjectDo {
	return c.Clauses(dbresolver.Write)
}

func (c cmsSubjectDo) Clauses(conds ...clause.Expression) *cmsSubjectDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cmsSubjectDo) Returning(value interface{}, columns ...string) *cmsSubjectDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cmsSubjectDo) Not(conds ...gen.Condition) *cmsSubjectDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cmsSubjectDo) Or(conds ...gen.Condition) *cmsSubjectDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cmsSubjectDo) Select(conds ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cmsSubjectDo) Where(conds ...gen.Condition) *cmsSubjectDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cmsSubjectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *cmsSubjectDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cmsSubjectDo) Order(conds ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cmsSubjectDo) Distinct(cols ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cmsSubjectDo) Omit(cols ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cmsSubjectDo) Join(table schema.Tabler, on ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cmsSubjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cmsSubjectDo) RightJoin(table schema.Tabler, on ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cmsSubjectDo) Group(cols ...field.Expr) *cmsSubjectDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cmsSubjectDo) Having(conds ...gen.Condition) *cmsSubjectDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cmsSubjectDo) Limit(limit int) *cmsSubjectDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cmsSubjectDo) Offset(offset int) *cmsSubjectDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cmsSubjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cmsSubjectDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cmsSubjectDo) Unscoped() *cmsSubjectDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cmsSubjectDo) Create(values ...*model.CmsSubject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cmsSubjectDo) CreateInBatches(values []*model.CmsSubject, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cmsSubjectDo) Save(values ...*model.CmsSubject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cmsSubjectDo) First() (*model.CmsSubject, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubject), nil
	}
}

func (c cmsSubjectDo) Take() (*model.CmsSubject, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubject), nil
	}
}

func (c cmsSubjectDo) Last() (*model.CmsSubject, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubject), nil
	}
}

func (c cmsSubjectDo) Find() ([]*model.CmsSubject, error) {
	result, err := c.DO.Find()
	return result.([]*model.CmsSubject), err
}

func (c cmsSubjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsSubject, err error) {
	buf := make([]*model.CmsSubject, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cmsSubjectDo) FindInBatches(result *[]*model.CmsSubject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cmsSubjectDo) Attrs(attrs ...field.AssignExpr) *cmsSubjectDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cmsSubjectDo) Assign(attrs ...field.AssignExpr) *cmsSubjectDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cmsSubjectDo) Joins(fields ...field.RelationField) *cmsSubjectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cmsSubjectDo) Preload(fields ...field.RelationField) *cmsSubjectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cmsSubjectDo) FirstOrInit() (*model.CmsSubject, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubject), nil
	}
}

func (c cmsSubjectDo) FirstOrCreate() (*model.CmsSubject, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubject), nil
	}
}

func (c cmsSubjectDo) FindByPage(offset int, limit int) (result []*model.CmsSubject, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cmsSubjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cmsSubjectDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cmsSubjectDo) Delete(models ...*model.CmsSubject) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cmsSubjectDo) withDO(do gen.Dao) *cmsSubjectDo {
	c.DO = *do.(*gen.DO)
	return c
}