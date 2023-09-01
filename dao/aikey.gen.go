// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"chatgpt-web-new-go/model"
)

func newAikey(db *gorm.DB, opts ...gen.DOOption) aikey {
	_aikey := aikey{}

	_aikey.aikeyDo.UseDB(db, opts...)
	_aikey.aikeyDo.UseModel(&model.Aikey{})

	tableName := _aikey.aikeyDo.TableName()
	_aikey.ALL = field.NewAsterisk(tableName)
	_aikey.ID = field.NewInt64(tableName, "id")
	_aikey.Key = field.NewString(tableName, "key")
	_aikey.Host = field.NewString(tableName, "host")
	_aikey.Remarks = field.NewString(tableName, "remarks")
	_aikey.Type = field.NewString(tableName, "type")
	_aikey.Models = field.NewString(tableName, "models")
	_aikey.Check = field.NewInt32(tableName, "check")
	_aikey.Limit_ = field.NewFloat64(tableName, "limit")
	_aikey.Usage = field.NewFloat64(tableName, "usage")
	_aikey.Status = field.NewInt32(tableName, "status")
	_aikey.CreateTime = field.NewTime(tableName, "create_time")
	_aikey.UpdateTime = field.NewTime(tableName, "update_time")
	_aikey.IsDelete = field.NewInt32(tableName, "is_delete")

	_aikey.fillFieldMap()

	return _aikey
}

type aikey struct {
	aikeyDo

	ALL        field.Asterisk
	ID         field.Int64
	Key        field.String
	Host       field.String
	Remarks    field.String
	Type       field.String  // openai sd
	Models     field.String  // 可用模型
	Check      field.Int32   // 1 检查token可用性 0不检查
	Limit_     field.Float64 // 总限制
	Usage      field.Float64 // 已经使用
	Status     field.Int32   // 1 正常 0异常
	CreateTime field.Time
	UpdateTime field.Time
	IsDelete   field.Int32

	fieldMap map[string]field.Expr
}

func (a aikey) Table(newTableName string) *aikey {
	a.aikeyDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aikey) As(alias string) *aikey {
	a.aikeyDo.DO = *(a.aikeyDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aikey) updateTableName(table string) *aikey {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Key = field.NewString(table, "key")
	a.Host = field.NewString(table, "host")
	a.Remarks = field.NewString(table, "remarks")
	a.Type = field.NewString(table, "type")
	a.Models = field.NewString(table, "models")
	a.Check = field.NewInt32(table, "check")
	a.Limit_ = field.NewFloat64(table, "limit")
	a.Usage = field.NewFloat64(table, "usage")
	a.Status = field.NewInt32(table, "status")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.IsDelete = field.NewInt32(table, "is_delete")

	a.fillFieldMap()

	return a
}

func (a *aikey) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aikey) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 13)
	a.fieldMap["id"] = a.ID
	a.fieldMap["key"] = a.Key
	a.fieldMap["host"] = a.Host
	a.fieldMap["remarks"] = a.Remarks
	a.fieldMap["type"] = a.Type
	a.fieldMap["models"] = a.Models
	a.fieldMap["check"] = a.Check
	a.fieldMap["limit"] = a.Limit_
	a.fieldMap["usage"] = a.Usage
	a.fieldMap["status"] = a.Status
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["is_delete"] = a.IsDelete
}

func (a aikey) clone(db *gorm.DB) aikey {
	a.aikeyDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aikey) replaceDB(db *gorm.DB) aikey {
	a.aikeyDo.ReplaceDB(db)
	return a
}

type aikeyDo struct{ gen.DO }

type IAikeyDo interface {
	gen.SubQuery
	Debug() IAikeyDo
	WithContext(ctx context.Context) IAikeyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAikeyDo
	WriteDB() IAikeyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAikeyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAikeyDo
	Not(conds ...gen.Condition) IAikeyDo
	Or(conds ...gen.Condition) IAikeyDo
	Select(conds ...field.Expr) IAikeyDo
	Where(conds ...gen.Condition) IAikeyDo
	Order(conds ...field.Expr) IAikeyDo
	Distinct(cols ...field.Expr) IAikeyDo
	Omit(cols ...field.Expr) IAikeyDo
	Join(table schema.Tabler, on ...field.Expr) IAikeyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAikeyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAikeyDo
	Group(cols ...field.Expr) IAikeyDo
	Having(conds ...gen.Condition) IAikeyDo
	Limit(limit int) IAikeyDo
	Offset(offset int) IAikeyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAikeyDo
	Unscoped() IAikeyDo
	Create(values ...*model.Aikey) error
	CreateInBatches(values []*model.Aikey, batchSize int) error
	Save(values ...*model.Aikey) error
	First() (*model.Aikey, error)
	Take() (*model.Aikey, error)
	Last() (*model.Aikey, error)
	Find() ([]*model.Aikey, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Aikey, err error)
	FindInBatches(result *[]*model.Aikey, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Aikey) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAikeyDo
	Assign(attrs ...field.AssignExpr) IAikeyDo
	Joins(fields ...field.RelationField) IAikeyDo
	Preload(fields ...field.RelationField) IAikeyDo
	FirstOrInit() (*model.Aikey, error)
	FirstOrCreate() (*model.Aikey, error)
	FindByPage(offset int, limit int) (result []*model.Aikey, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAikeyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Aikey, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (a aikeyDo) FilterWithNameAndRole(name string, role string) (result []model.Aikey, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM aikey WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a aikeyDo) Debug() IAikeyDo {
	return a.withDO(a.DO.Debug())
}

func (a aikeyDo) WithContext(ctx context.Context) IAikeyDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aikeyDo) ReadDB() IAikeyDo {
	return a.Clauses(dbresolver.Read)
}

func (a aikeyDo) WriteDB() IAikeyDo {
	return a.Clauses(dbresolver.Write)
}

func (a aikeyDo) Session(config *gorm.Session) IAikeyDo {
	return a.withDO(a.DO.Session(config))
}

func (a aikeyDo) Clauses(conds ...clause.Expression) IAikeyDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aikeyDo) Returning(value interface{}, columns ...string) IAikeyDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aikeyDo) Not(conds ...gen.Condition) IAikeyDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aikeyDo) Or(conds ...gen.Condition) IAikeyDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aikeyDo) Select(conds ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aikeyDo) Where(conds ...gen.Condition) IAikeyDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aikeyDo) Order(conds ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aikeyDo) Distinct(cols ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aikeyDo) Omit(cols ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aikeyDo) Join(table schema.Tabler, on ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aikeyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aikeyDo) RightJoin(table schema.Tabler, on ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aikeyDo) Group(cols ...field.Expr) IAikeyDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aikeyDo) Having(conds ...gen.Condition) IAikeyDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aikeyDo) Limit(limit int) IAikeyDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aikeyDo) Offset(offset int) IAikeyDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aikeyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAikeyDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aikeyDo) Unscoped() IAikeyDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aikeyDo) Create(values ...*model.Aikey) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aikeyDo) CreateInBatches(values []*model.Aikey, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aikeyDo) Save(values ...*model.Aikey) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aikeyDo) First() (*model.Aikey, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Aikey), nil
	}
}

func (a aikeyDo) Take() (*model.Aikey, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Aikey), nil
	}
}

func (a aikeyDo) Last() (*model.Aikey, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Aikey), nil
	}
}

func (a aikeyDo) Find() ([]*model.Aikey, error) {
	result, err := a.DO.Find()
	return result.([]*model.Aikey), err
}

func (a aikeyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Aikey, err error) {
	buf := make([]*model.Aikey, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aikeyDo) FindInBatches(result *[]*model.Aikey, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aikeyDo) Attrs(attrs ...field.AssignExpr) IAikeyDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aikeyDo) Assign(attrs ...field.AssignExpr) IAikeyDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aikeyDo) Joins(fields ...field.RelationField) IAikeyDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aikeyDo) Preload(fields ...field.RelationField) IAikeyDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aikeyDo) FirstOrInit() (*model.Aikey, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Aikey), nil
	}
}

func (a aikeyDo) FirstOrCreate() (*model.Aikey, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Aikey), nil
	}
}

func (a aikeyDo) FindByPage(offset int, limit int) (result []*model.Aikey, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a aikeyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aikeyDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aikeyDo) Delete(models ...*model.Aikey) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aikeyDo) withDO(do gen.Dao) *aikeyDo {
	a.DO = *do.(*gen.DO)
	return a
}
