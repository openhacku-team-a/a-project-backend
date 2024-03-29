// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gQuery

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"a-project-backend/gen/gModel"
)

func newUserTag(db *gorm.DB, opts ...gen.DOOption) userTag {
	_userTag := userTag{}

	_userTag.userTagDo.UseDB(db, opts...)
	_userTag.userTagDo.UseModel(&gModel.UserTag{})

	tableName := _userTag.userTagDo.TableName()
	_userTag.ALL = field.NewAsterisk(tableName)
	_userTag.UserID = field.NewString(tableName, "user_id")
	_userTag.TagID = field.NewString(tableName, "tag_id")

	_userTag.fillFieldMap()

	return _userTag
}

type userTag struct {
	userTagDo

	ALL    field.Asterisk
	UserID field.String
	TagID  field.String

	fieldMap map[string]field.Expr
}

func (u userTag) Table(newTableName string) *userTag {
	u.userTagDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userTag) As(alias string) *userTag {
	u.userTagDo.DO = *(u.userTagDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userTag) updateTableName(table string) *userTag {
	u.ALL = field.NewAsterisk(table)
	u.UserID = field.NewString(table, "user_id")
	u.TagID = field.NewString(table, "tag_id")

	u.fillFieldMap()

	return u
}

func (u *userTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userTag) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 2)
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["tag_id"] = u.TagID
}

func (u userTag) clone(db *gorm.DB) userTag {
	u.userTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userTag) replaceDB(db *gorm.DB) userTag {
	u.userTagDo.ReplaceDB(db)
	return u
}

type userTagDo struct{ gen.DO }

type IUserTagDo interface {
	gen.SubQuery
	Debug() IUserTagDo
	WithContext(ctx context.Context) IUserTagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserTagDo
	WriteDB() IUserTagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserTagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserTagDo
	Not(conds ...gen.Condition) IUserTagDo
	Or(conds ...gen.Condition) IUserTagDo
	Select(conds ...field.Expr) IUserTagDo
	Where(conds ...gen.Condition) IUserTagDo
	Order(conds ...field.Expr) IUserTagDo
	Distinct(cols ...field.Expr) IUserTagDo
	Omit(cols ...field.Expr) IUserTagDo
	Join(table schema.Tabler, on ...field.Expr) IUserTagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserTagDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserTagDo
	Group(cols ...field.Expr) IUserTagDo
	Having(conds ...gen.Condition) IUserTagDo
	Limit(limit int) IUserTagDo
	Offset(offset int) IUserTagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTagDo
	Unscoped() IUserTagDo
	Create(values ...*gModel.UserTag) error
	CreateInBatches(values []*gModel.UserTag, batchSize int) error
	Save(values ...*gModel.UserTag) error
	First() (*gModel.UserTag, error)
	Take() (*gModel.UserTag, error)
	Last() (*gModel.UserTag, error)
	Find() ([]*gModel.UserTag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gModel.UserTag, err error)
	FindInBatches(result *[]*gModel.UserTag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*gModel.UserTag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserTagDo
	Assign(attrs ...field.AssignExpr) IUserTagDo
	Joins(fields ...field.RelationField) IUserTagDo
	Preload(fields ...field.RelationField) IUserTagDo
	FirstOrInit() (*gModel.UserTag, error)
	FirstOrCreate() (*gModel.UserTag, error)
	FindByPage(offset int, limit int) (result []*gModel.UserTag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserTagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userTagDo) Debug() IUserTagDo {
	return u.withDO(u.DO.Debug())
}

func (u userTagDo) WithContext(ctx context.Context) IUserTagDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userTagDo) ReadDB() IUserTagDo {
	return u.Clauses(dbresolver.Read)
}

func (u userTagDo) WriteDB() IUserTagDo {
	return u.Clauses(dbresolver.Write)
}

func (u userTagDo) Session(config *gorm.Session) IUserTagDo {
	return u.withDO(u.DO.Session(config))
}

func (u userTagDo) Clauses(conds ...clause.Expression) IUserTagDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userTagDo) Returning(value interface{}, columns ...string) IUserTagDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userTagDo) Not(conds ...gen.Condition) IUserTagDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userTagDo) Or(conds ...gen.Condition) IUserTagDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userTagDo) Select(conds ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userTagDo) Where(conds ...gen.Condition) IUserTagDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userTagDo) Order(conds ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userTagDo) Distinct(cols ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userTagDo) Omit(cols ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userTagDo) Join(table schema.Tabler, on ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userTagDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userTagDo) Group(cols ...field.Expr) IUserTagDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userTagDo) Having(conds ...gen.Condition) IUserTagDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userTagDo) Limit(limit int) IUserTagDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userTagDo) Offset(offset int) IUserTagDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTagDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userTagDo) Unscoped() IUserTagDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userTagDo) Create(values ...*gModel.UserTag) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userTagDo) CreateInBatches(values []*gModel.UserTag, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userTagDo) Save(values ...*gModel.UserTag) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userTagDo) First() (*gModel.UserTag, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gModel.UserTag), nil
	}
}

func (u userTagDo) Take() (*gModel.UserTag, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gModel.UserTag), nil
	}
}

func (u userTagDo) Last() (*gModel.UserTag, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gModel.UserTag), nil
	}
}

func (u userTagDo) Find() ([]*gModel.UserTag, error) {
	result, err := u.DO.Find()
	return result.([]*gModel.UserTag), err
}

func (u userTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gModel.UserTag, err error) {
	buf := make([]*gModel.UserTag, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userTagDo) FindInBatches(result *[]*gModel.UserTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userTagDo) Attrs(attrs ...field.AssignExpr) IUserTagDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userTagDo) Assign(attrs ...field.AssignExpr) IUserTagDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userTagDo) Joins(fields ...field.RelationField) IUserTagDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userTagDo) Preload(fields ...field.RelationField) IUserTagDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userTagDo) FirstOrInit() (*gModel.UserTag, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gModel.UserTag), nil
	}
}

func (u userTagDo) FirstOrCreate() (*gModel.UserTag, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gModel.UserTag), nil
	}
}

func (u userTagDo) FindByPage(offset int, limit int) (result []*gModel.UserTag, count int64, err error) {
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

func (u userTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userTagDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userTagDo) Delete(models ...*gModel.UserTag) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userTagDo) withDO(do gen.Dao) *userTagDo {
	u.DO = *do.(*gen.DO)
	return u
}
