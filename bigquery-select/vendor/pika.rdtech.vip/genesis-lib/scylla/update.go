package scylla

import (
	"time"

	"pika.rdtech.vip/eden-lib/gocqlx/v2/qb"
)

type UpdateBuilder struct {
	qb   *qb.UpdateBuilder
	cqlx *Cqlx
}

// type op byte

// type Cmp struct {
// 	op     op
// 	column string
// 	value  value
// }

// type value interface {
// 	writeCql(cql *bytes.Buffer) (names []string)
// }

// type param string

// func (p param) writeCql(cql *bytes.Buffer) (names []string) {
// 	cql.WriteByte('?')
// 	return []string{string(p)}
// }

func (b *UpdateBuilder) Query() *Queryx {
	stmt, names := b.qb.ToCql()

	return &Queryx{
		b.cqlx.Session.Query(stmt, names),
	}
}

func (b *UpdateBuilder) ToCql() (stmt string, names []string) {
	return b.qb.ToCql()
}

// Table sets the table to be updated.
func (b *UpdateBuilder) Table(table string) *UpdateBuilder {
	b.qb.Table(table)
	return b
}

func (b *UpdateBuilder) TTL(d time.Duration) *UpdateBuilder {
	b.qb.TTL(d)
	return b
}

func (b *UpdateBuilder) TTLNamed(name string) *UpdateBuilder {
	b.qb.TTLNamed(name)
	return b
}

func (b *UpdateBuilder) Timestamp(t time.Time) *UpdateBuilder {
	b.qb.Timestamp(t)
	return b
}

func (b *UpdateBuilder) TimestampNamed(name string) *UpdateBuilder {
	b.qb.TimestampNamed(name)
	return b
}

func (b *UpdateBuilder) Set(columns ...string) *UpdateBuilder {
	b.qb.Set(columns...)
	return b
}

func (b *UpdateBuilder) SetNamed(column, name string) *UpdateBuilder {
	b.qb.SetNamed(column, name)
	return b
}

func (b *UpdateBuilder) SetLit(column, literal string) *UpdateBuilder {
	b.qb.SetLit(column, literal)
	return b
}

func (b *UpdateBuilder) SetFunc(column string, ff *Func) *UpdateBuilder {
	fn := &qb.Func{
		Name:       ff.Name,
		ParamNames: ff.ParamNames,
	}
	b.qb.SetFunc(column, fn)
	return b
}

func (b *UpdateBuilder) SetTuple(column string, count int) *UpdateBuilder {
	b.qb.SetTuple(column, count)
	return b
}

func (b *UpdateBuilder) Add(column string) *UpdateBuilder {
	b.qb.Add(column)
	return b
}

func (b *UpdateBuilder) AddNamed(column, name string) *UpdateBuilder {
	b.qb.AddNamed(column, name)
	return b
}

func (b *UpdateBuilder) AddLit(column, literal string) *UpdateBuilder {
	b.qb.AddLit(column, literal)
	return b
}

func (b *UpdateBuilder) AddFunc(column string, ff *Func) *UpdateBuilder {
	fn := &qb.Func{
		Name:       ff.Name,
		ParamNames: ff.ParamNames,
	}
	b.qb.AddFunc(column, fn)
	return b
}

func (b *UpdateBuilder) Remove(column string) *UpdateBuilder {
	b.qb.Remove(column)
	return b
}

func (b *UpdateBuilder) RemoveNamed(column, name string) *UpdateBuilder {
	b.qb.RemoveNamed(column, name)
	return b
}

func (b *UpdateBuilder) RemoveLit(column, literal string) *UpdateBuilder {
	b.qb.RemoveLit(column, literal)
	return b
}

func (b *UpdateBuilder) RemoveFunc(column string, ff *Func) *UpdateBuilder {
	fn := &qb.Func{
		Name:       ff.Name,
		ParamNames: ff.ParamNames,
	}
	b.qb.RemoveFunc(column, fn)

	return b
}

func (b *UpdateBuilder) Where(w ...qb.Cmp) *UpdateBuilder {
	// for _, v := range w {
	b.qb.Where(w...)
	// }

	return b
}

func (b *UpdateBuilder) WhereTuple(w []string, count int) *UpdateBuilder {
	for _, v := range w {
		b.qb.Where(qb.EqTuple(v, count))
	}

	return b
}

func (b *UpdateBuilder) WhereqNamed(w []string, name string) *UpdateBuilder {
	for _, v := range w {
		b.qb.Where(qb.EqNamed(v, name))
	}

	return b
}

func (b *UpdateBuilder) If(w ...qb.Cmp) *UpdateBuilder {
	// for _, v := range w {
	b.qb.If(w...)
	// }
	return b
}

func (b *UpdateBuilder) Existing() *UpdateBuilder {
	b.qb.Existing()
	return b
}
