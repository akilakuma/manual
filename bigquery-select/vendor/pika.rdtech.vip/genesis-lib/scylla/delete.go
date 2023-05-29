package scylla

import (
	"pika.rdtech.vip/eden-lib/gocqlx/v2/qb"
)

type DeleteBuilder struct {
	qb   *qb.DeleteBuilder
	cqlx *Cqlx
}

func (b *DeleteBuilder) Query() *Queryx {
	stmt, names := b.qb.ToCql()

	return &Queryx{
		b.cqlx.Session.Query(stmt, names),
	}
}

func (b *DeleteBuilder) ToCql() (stmt string, names []string) {
	return b.qb.ToCql()
}

func (b *DeleteBuilder) Where(w ...qb.Cmp) *DeleteBuilder {
	// for _, v := range w {
	b.qb.Where(w...)
	// }

	return b
}

func (b *DeleteBuilder) WhereTuple(w []string, count int) *DeleteBuilder {
	for _, v := range w {
		b.qb.Where(qb.EqTuple(v, count))
	}

	return b
}

func (b *DeleteBuilder) WhereqNamed(w []string, name string) *DeleteBuilder {
	for _, v := range w {
		b.qb.Where(qb.EqNamed(v, name))
	}

	return b
}

func (b *DeleteBuilder) If(w ...qb.Cmp) *DeleteBuilder {
	// for _, v := range w {
	b.qb.If(w...)
	// }
	return b
}

func (b *DeleteBuilder) Existing() *DeleteBuilder {
	b.qb.Existing()
	return b
}
