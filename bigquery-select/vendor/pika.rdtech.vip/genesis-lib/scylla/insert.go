package scylla

import (
	"time"

	"pika.rdtech.vip/eden-lib/gocqlx/v2/qb"
)

type InsertBuilder struct {
	qbi  *qb.InsertBuilder
	cqlx *Cqlx
}

func (i *InsertBuilder) Query() *Queryx {
	stmt, names := i.qbi.ToCql()

	return &Queryx{
		i.cqlx.Session.Query(stmt, names),
	}
}

func (i *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	i.qbi.Columns(columns...)

	return i
}

func (i *InsertBuilder) FuncColumn(column string, fn *Func) *InsertBuilder {
	i.qbi.FuncColumn(column, &qb.Func{
		Name:       fn.Name,
		ParamNames: fn.ParamNames,
	})

	return i
}

func (i *InsertBuilder) Into(table string) *InsertBuilder {
	i.qbi.Into(table)

	return i
}

func (i *InsertBuilder) Json() *InsertBuilder {
	i.qbi.Json()

	return i
}

func (i *InsertBuilder) LitColumn(column, literal string) *InsertBuilder {
	i.qbi.LitColumn(column, literal)

	return i
}

func (i *InsertBuilder) NamedColumn(column, name string) *InsertBuilder {
	i.qbi.NamedColumn(column, name)

	return i
}

func (i *InsertBuilder) TTL(d time.Duration) *InsertBuilder {
	i.qbi.TTL(d)

	return i
}

func (i *InsertBuilder) TTLNamed(name string) *InsertBuilder {
	i.qbi.TTLNamed(name)

	return i
}

func (i *InsertBuilder) Timestamp(t time.Time) *InsertBuilder {
	i.qbi.Timestamp(t)

	return i
}

func (i *InsertBuilder) TimestampNamed(name string) *InsertBuilder {
	i.qbi.TimestampNamed(name)

	return i
}

func (i *InsertBuilder) ToCql() (stmt string, names []string) {
	return i.qbi.ToCql()
}

func (i *InsertBuilder) TupleColumn(column string, count int) *InsertBuilder {
	i.qbi.TupleColumn(column, count)

	return i
}

func (i *InsertBuilder) Unique() *InsertBuilder {
	i.qbi.Unique()

	return i
}

type TableInsertBuilder struct {
	qbi *qb.InsertBuilder
}

func (ti *TableInsertBuilder) Columns(columns ...string) *TableInsertBuilder {
	ti.qbi.Columns(columns...)

	return ti
}

func (ti *TableInsertBuilder) FuncColumn(column string, fn *Func) *TableInsertBuilder {
	ti.qbi.FuncColumn(column, &qb.Func{
		Name:       fn.Name,
		ParamNames: fn.ParamNames,
	})

	return ti
}

func (ti *TableInsertBuilder) Into(table string) *TableInsertBuilder {
	ti.qbi.Into(table)

	return ti
}

func (ti *TableInsertBuilder) Json() *TableInsertBuilder {
	ti.qbi.Json()

	return ti
}

func (ti *TableInsertBuilder) LitColumn(column, literal string) *TableInsertBuilder {
	ti.qbi.LitColumn(column, literal)

	return ti
}

func (ti *TableInsertBuilder) NamedColumn(column, name string) *TableInsertBuilder {
	ti.qbi.NamedColumn(column, name)

	return ti
}

func (ti *TableInsertBuilder) TTL(d time.Duration) *TableInsertBuilder {
	ti.qbi.TTL(d)

	return ti
}

func (ti *TableInsertBuilder) TTLNamed(name string) *TableInsertBuilder {
	ti.qbi.TTLNamed(name)

	return ti
}

func (ti *TableInsertBuilder) Timestamp(t time.Time) *TableInsertBuilder {
	ti.qbi.Timestamp(t)

	return ti
}

func (ti *TableInsertBuilder) TimestampNamed(name string) *TableInsertBuilder {
	ti.qbi.TimestampNamed(name)

	return ti
}

func (ti *TableInsertBuilder) ToCql() (stmt string, names []string) {
	return ti.qbi.ToCql()
}

func (ti *TableInsertBuilder) TupleColumn(column string, count int) *TableInsertBuilder {
	ti.qbi.TupleColumn(column, count)

	return ti
}

func (ti *TableInsertBuilder) Unique() *TableInsertBuilder {
	ti.qbi.Unique()

	return ti
}
