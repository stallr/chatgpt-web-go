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

func newPlugin(db *gorm.DB, opts ...gen.DOOption) plugin {
	_plugin := plugin{}

	_plugin.pluginDo.UseDB(db, opts...)
	_plugin.pluginDo.UseModel(&model.Plugin{})

	tableName := _plugin.pluginDo.TableName()
	_plugin.ALL = field.NewAsterisk(tableName)
	_plugin.ID = field.NewInt64(tableName, "id")
	_plugin.Name = field.NewString(tableName, "name")
	_plugin.Description = field.NewString(tableName, "description")
	_plugin.Avatar = field.NewString(tableName, "avatar")
	_plugin.Variables = field.NewString(tableName, "variables")
	_plugin.Function = field.NewString(tableName, "function")
	_plugin.Script = field.NewString(tableName, "script")
	_plugin.UserID = field.NewInt64(tableName, "user_id")
	_plugin.Status = field.NewInt32(tableName, "status")
	_plugin.CreateTime = field.NewTime(tableName, "create_time")
	_plugin.UpdateTime = field.NewTime(tableName, "update_time")
	_plugin.IsDelete = field.NewInt32(tableName, "is_delete")

	_plugin.fillFieldMap()

	return _plugin
}

type plugin struct {
	pluginDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String // 插件名称
	Description field.String // 插件描述
	Avatar      field.String // 插件头像
	Variables   field.String // 变量
	Function    field.String // 函数配置文件
	Script      field.String // js脚本
	UserID      field.Int64  // 提交用户id
	Status      field.Int32  // 4为审核中1为正常0为异常
	CreateTime  field.Time
	UpdateTime  field.Time
	IsDelete    field.Int32

	fieldMap map[string]field.Expr
}

func (p plugin) Table(newTableName string) *plugin {
	p.pluginDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p plugin) As(alias string) *plugin {
	p.pluginDo.DO = *(p.pluginDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *plugin) updateTableName(table string) *plugin {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.Description = field.NewString(table, "description")
	p.Avatar = field.NewString(table, "avatar")
	p.Variables = field.NewString(table, "variables")
	p.Function = field.NewString(table, "function")
	p.Script = field.NewString(table, "script")
	p.UserID = field.NewInt64(table, "user_id")
	p.Status = field.NewInt32(table, "status")
	p.CreateTime = field.NewTime(table, "create_time")
	p.UpdateTime = field.NewTime(table, "update_time")
	p.IsDelete = field.NewInt32(table, "is_delete")

	p.fillFieldMap()

	return p
}

func (p *plugin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *plugin) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["description"] = p.Description
	p.fieldMap["avatar"] = p.Avatar
	p.fieldMap["variables"] = p.Variables
	p.fieldMap["function"] = p.Function
	p.fieldMap["script"] = p.Script
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["status"] = p.Status
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["update_time"] = p.UpdateTime
	p.fieldMap["is_delete"] = p.IsDelete
}

func (p plugin) clone(db *gorm.DB) plugin {
	p.pluginDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p plugin) replaceDB(db *gorm.DB) plugin {
	p.pluginDo.ReplaceDB(db)
	return p
}

type pluginDo struct{ gen.DO }

type IPluginDo interface {
	gen.SubQuery
	Debug() IPluginDo
	WithContext(ctx context.Context) IPluginDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPluginDo
	WriteDB() IPluginDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPluginDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPluginDo
	Not(conds ...gen.Condition) IPluginDo
	Or(conds ...gen.Condition) IPluginDo
	Select(conds ...field.Expr) IPluginDo
	Where(conds ...gen.Condition) IPluginDo
	Order(conds ...field.Expr) IPluginDo
	Distinct(cols ...field.Expr) IPluginDo
	Omit(cols ...field.Expr) IPluginDo
	Join(table schema.Tabler, on ...field.Expr) IPluginDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPluginDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPluginDo
	Group(cols ...field.Expr) IPluginDo
	Having(conds ...gen.Condition) IPluginDo
	Limit(limit int) IPluginDo
	Offset(offset int) IPluginDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPluginDo
	Unscoped() IPluginDo
	Create(values ...*model.Plugin) error
	CreateInBatches(values []*model.Plugin, batchSize int) error
	Save(values ...*model.Plugin) error
	First() (*model.Plugin, error)
	Take() (*model.Plugin, error)
	Last() (*model.Plugin, error)
	Find() ([]*model.Plugin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Plugin, err error)
	FindInBatches(result *[]*model.Plugin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Plugin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPluginDo
	Assign(attrs ...field.AssignExpr) IPluginDo
	Joins(fields ...field.RelationField) IPluginDo
	Preload(fields ...field.RelationField) IPluginDo
	FirstOrInit() (*model.Plugin, error)
	FirstOrCreate() (*model.Plugin, error)
	FindByPage(offset int, limit int) (result []*model.Plugin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPluginDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Plugin, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (p pluginDo) FilterWithNameAndRole(name string, role string) (result []model.Plugin, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM plugin WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p pluginDo) Debug() IPluginDo {
	return p.withDO(p.DO.Debug())
}

func (p pluginDo) WithContext(ctx context.Context) IPluginDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pluginDo) ReadDB() IPluginDo {
	return p.Clauses(dbresolver.Read)
}

func (p pluginDo) WriteDB() IPluginDo {
	return p.Clauses(dbresolver.Write)
}

func (p pluginDo) Session(config *gorm.Session) IPluginDo {
	return p.withDO(p.DO.Session(config))
}

func (p pluginDo) Clauses(conds ...clause.Expression) IPluginDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pluginDo) Returning(value interface{}, columns ...string) IPluginDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pluginDo) Not(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pluginDo) Or(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pluginDo) Select(conds ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pluginDo) Where(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pluginDo) Order(conds ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pluginDo) Distinct(cols ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pluginDo) Omit(cols ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pluginDo) Join(table schema.Tabler, on ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pluginDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPluginDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pluginDo) RightJoin(table schema.Tabler, on ...field.Expr) IPluginDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pluginDo) Group(cols ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pluginDo) Having(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pluginDo) Limit(limit int) IPluginDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pluginDo) Offset(offset int) IPluginDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pluginDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPluginDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pluginDo) Unscoped() IPluginDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pluginDo) Create(values ...*model.Plugin) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pluginDo) CreateInBatches(values []*model.Plugin, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pluginDo) Save(values ...*model.Plugin) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pluginDo) First() (*model.Plugin, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plugin), nil
	}
}

func (p pluginDo) Take() (*model.Plugin, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plugin), nil
	}
}

func (p pluginDo) Last() (*model.Plugin, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plugin), nil
	}
}

func (p pluginDo) Find() ([]*model.Plugin, error) {
	result, err := p.DO.Find()
	return result.([]*model.Plugin), err
}

func (p pluginDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Plugin, err error) {
	buf := make([]*model.Plugin, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pluginDo) FindInBatches(result *[]*model.Plugin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pluginDo) Attrs(attrs ...field.AssignExpr) IPluginDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pluginDo) Assign(attrs ...field.AssignExpr) IPluginDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pluginDo) Joins(fields ...field.RelationField) IPluginDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pluginDo) Preload(fields ...field.RelationField) IPluginDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pluginDo) FirstOrInit() (*model.Plugin, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plugin), nil
	}
}

func (p pluginDo) FirstOrCreate() (*model.Plugin, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plugin), nil
	}
}

func (p pluginDo) FindByPage(offset int, limit int) (result []*model.Plugin, count int64, err error) {
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

func (p pluginDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pluginDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pluginDo) Delete(models ...*model.Plugin) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pluginDo) withDO(do gen.Dao) *pluginDo {
	p.DO = *do.(*gen.DO)
	return p
}
