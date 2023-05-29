package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

func (o *Orm) Create(value interface{}) *Orm {
	return &Orm{db: o.db.Create(value)}
}

func (o *Orm) Find(value interface{}) *Orm {
	return &Orm{db: o.db.Find(value)}
}

func (o *Orm) Model(value interface{}) *Orm {
	return &Orm{db: o.db.Model(value)}
}

// 指定回傳值 如果沒撈到也會返回預設值
func (o *Orm) Assign(value ...interface{}) *Orm {
	return &Orm{db: o.db.Assign(value...)}
}

// 指定回傳值 如果沒撈到也會返回預設值
func (o *Orm) First(out interface{}, where ...interface{}) *Orm {
	return &Orm{db: o.db.First(out, where...)}
}

func (o *Orm) FirstOrInit(out interface{}, where ...interface{}) *Orm {
	return &Orm{db: o.db.FirstOrInit(out, where...)}
}

func (o *Orm) Attrs(attr ...interface{}) *Orm {
	return &Orm{db: o.db.Attrs(attr...)}
}

func (o *Orm) CreateTable(models ...interface{}) *Orm {
	return &Orm{db: o.db.CreateTable(models...)}
}

func (o *Orm) Debug() *Orm {
	return &Orm{db: o.db.Debug()}
}

func (o *Orm) Begin() *Orm {
	return &Orm{db: o.db.Begin()}
}

func (o *Orm) BeginTx(ctx context.Context, opts sql.TxOptions) *Orm {
	return &Orm{db: o.db.BeginTx(ctx, &opts)}
}

func (o *Orm) Commit() *Orm {
	return &Orm{db: o.db.Commit()}
}

func (o *Orm) CommonDB() SQLCommon {
	return o.db.CommonDB()
}

func (o *Orm) Count(value interface{}) *Orm {
	return &Orm{db: o.db.Count(value)}
}

func (o *Orm) Delete(value interface{}, where ...interface{}) *Orm {
	// 這是真的刪除
	return &Orm{db: o.db.Unscoped().Delete(value, where...)}
}

func (o *Orm) Exec(sql string, values ...interface{}) *Orm {
	return &Orm{db: o.db.Exec(sql, values...)}
}

func (o *Orm) FirstOrCreate(out interface{}, where ...interface{}) *Orm {
	return &Orm{db: o.db.FirstOrCreate(out, where...)}
}

func (o *Orm) Get(name string) (value interface{}, ok bool) {
	return o.db.Get(name)
}

func (o *Orm) GetErrors() []error {
	return o.db.GetErrors()
}

func (o *Orm) Group(query string) *Orm {
	return &Orm{db: o.db.Group(query)}
}

func (o *Orm) HasBlockGlobalUpdate() bool {
	return o.db.HasBlockGlobalUpdate()
}

func (o *Orm) HasTable(value interface{}) bool {
	return o.db.HasTable(value)
}

func (o *Orm) Having(query interface{}, values ...interface{}) *Orm {
	return &Orm{db: o.db.Having(query, values...)}
}

func (o *Orm) InstantSet(name string, value interface{}) *Orm {
	return &Orm{db: o.db.InstantSet(name, value)}
}

func (o *Orm) Joins(query string, args ...interface{}) *Orm {
	return &Orm{db: o.db.Joins(query, args...)}
}

func (o *Orm) Last(out interface{}, where ...interface{}) *Orm {
	return &Orm{db: o.db.Last(out, where...)}
}

func (o *Orm) Limit(limit interface{}) *Orm {
	return &Orm{db: o.db.Limit(limit)}
}

func (o *Orm) LogMode(enable bool) *Orm {
	return &Orm{db: o.db.LogMode(enable)}
}

func (o *Orm) ModifyColumn(column string, typ string) *Orm {
	return &Orm{db: o.db.ModifyColumn(column, typ)}
}

func (o *Orm) New() *Orm {
	return &Orm{db: o.db.New()}
}

func (o *Orm) NewRecord(value interface{}) bool {
	return o.db.NewRecord(value)
}

func (o *Orm) NewScope(value interface{}) *Scope {
	s := o.db.NewScope(value)
	sc := &Scope{
		Scope: s,
	}

	return sc
}

func (o *Orm) Not(query interface{}, args ...interface{}) *Orm {
	return &Orm{db: o.db.Not(query, args...)}
}

func (o *Orm) Offset(offset interface{}) *Orm {
	return &Orm{db: o.db.Offset(offset)}
}

func (o *Orm) Omit(columns ...string) *Orm {
	return &Orm{db: o.db.Omit(columns...)}
}

func (o *Orm) Or(query interface{}, args ...interface{}) *Orm {
	return &Orm{db: o.db.Or(query, args...)}
}

func (o *Orm) Order(value interface{}, reorder ...bool) *Orm {
	return &Orm{db: o.db.Order(value, reorder...)}
}

func (o *Orm) Pluck(column string, value interface{}) *Orm {
	return &Orm{db: o.db.Pluck(column, value)}
}

func (o *Orm) Preload(column string, conditions ...interface{}) *Orm {
	return &Orm{db: o.db.Preload(column, conditions...)}
}

func (o *Orm) Preloads(out interface{}) *Orm {
	return &Orm{db: o.db.Preloads(out)}
}

func (o *Orm) QueryExpr() *SqlExpr {
	return &SqlExpr{SqlExpr: o.db.QueryExpr()}
}

func (o *Orm) Raw(sql string, values ...interface{}) *Orm {
	return &Orm{db: o.db.Raw(sql, values...)}
}

func (o *Orm) RecordNotFound() bool {
	return o.db.RecordNotFound()
}

func (o *Orm) Related(value interface{}, foreignKeys ...string) *Orm {
	return &Orm{db: o.db.Related(value, foreignKeys...)}
}

func (o *Orm) RemoveForeignKey(field string, dest string) *Orm {
	return &Orm{db: o.db.RemoveForeignKey(field, dest)}
}

func (o *Orm) RemoveIndex(indexName string) *Orm {
	return &Orm{db: o.db.RemoveIndex(indexName)}
}

func (o *Orm) Rollback() *Orm {
	return &Orm{db: o.db.Rollback()}
}

func (o *Orm) RollbackUnlessCommitted() *Orm {
	return &Orm{db: o.db.RollbackUnlessCommitted()}
}

func (o *Orm) Row() *sql.Row {
	return o.db.Row()
}

func (o *Orm) Rows() (*sql.Rows, error) {
	return o.db.Rows()
}

func (o *Orm) Save(value interface{}) *Orm {
	return &Orm{db: o.db.Save(value)}
}

func (o *Orm) Scan(dest interface{}) *Orm {
	return &Orm{db: o.db.Scan(dest)}
}

func (o *Orm) ScanRows(rows *sql.Rows, result interface{}) error {
	return o.db.ScanRows(rows, result)
}

func (o *Orm) Scopes(funcs ...func(*Orm) *Orm) *Orm {
	for _, f := range funcs {
		o = f(o)
	}

	return o
}

func (o *Orm) Select(query interface{}, args ...interface{}) *Orm {
	return &Orm{db: o.db.Select(query, args...)}
}

func (o *Orm) Set(name string, value interface{}) *Orm {
	return &Orm{db: o.db.Set(name, value)}
}

func (o *Orm) SetLogger(log logger) {
	o.db.SetLogger(log)
}

func (o *Orm) SetNowFuncOverride(nowFuncOverride func() time.Time) *Orm {
	return &Orm{db: o.db.SetNowFuncOverride(nowFuncOverride)}
}

func (o *Orm) SingularTable(enable bool) {
	o.db.SingularTable(enable)
}

func (o *Orm) SubQuery() *SqlExpr {
	return &SqlExpr{SqlExpr: o.db.SubQuery()}
}

func (o *Orm) Table(name string) *Orm {
	return &Orm{db: o.db.Table(name)}
}

func (o *Orm) Take(out interface{}, where ...interface{}) *Orm {
	return &Orm{db: o.db.Take(out, where...)}
}

func (o *Orm) Transaction(fc func(tx *Orm) error) (err error) {
	err = o.db.Transaction(func(tx *gorm.DB) error {
		inOrm := &Orm{db: tx}

		return fc(inOrm)
	})

	return err
}

func (o *Orm) Unscoped() *Orm {
	return &Orm{db: o.db.Unscoped()}
}

func (o *Orm) Update(attrs ...interface{}) *Orm {
	return &Orm{db: o.db.Update(transformSliceSqlExpr(attrs)...)}
}

func (o *Orm) UpdateColumn(attrs ...interface{}) *Orm {
	return &Orm{db: o.db.UpdateColumn(transformSliceSqlExpr(attrs)...)}
}

func (o *Orm) UpdateColumns(values interface{}) *Orm {
	return &Orm{db: o.db.UpdateColumns(transformMapSqlExpr(values))}
}

func (o *Orm) Updates(values interface{}, ignoreProtectedAttrs ...bool) *Orm {
	return &Orm{db: o.db.Updates(transformMapSqlExpr(values), ignoreProtectedAttrs...)}
}

func (o *Orm) Where(query interface{}, args ...interface{}) *Orm {
	return &Orm{db: o.db.Where(query, transformSliceSqlExpr(args)...)}
}

func (o *Orm) BlockGlobalUpdate(enable bool) *Orm {
	return &Orm{db: o.db.BlockGlobalUpdate(enable)}
}
