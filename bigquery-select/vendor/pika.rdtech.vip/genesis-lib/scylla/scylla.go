package scylla

import (
	"github.com/syhlion/gocql"

	"pika.rdtech.vip/eden-lib/gocqlx/v2"
	"pika.rdtech.vip/eden-lib/gocqlx/v2/qb"
)

type Cqlx struct {
	Session gocqlx.Session
}

func NewSession(hosts []string, opt ...Option) (*Cqlx, error) {
	opts := gocql.NewCluster(hosts...)

	for _, o := range opt {
		o.apply(opts)
	}

	s, err := gocqlx.WrapSession(opts.CreateSession())

	if err != nil {
		return nil, err
	}

	return &Cqlx{Session: s}, nil
}

func (c *Cqlx) Close() {
	c.Session.Close()
}

func (c *Cqlx) Insert(table string) *InsertBuilder {
	return &InsertBuilder{
		cqlx: c,
		qbi:  qb.Insert(table),
	}
}

func (c *Cqlx) Select(table string) *SelectBuilder {
	return &SelectBuilder{
		cqlx: c,
		qbs:  qb.Select(table),
	}
}

func (c *Cqlx) Update(table string) *UpdateBuilder {
	return &UpdateBuilder{
		cqlx: c,
		qb:   qb.Update(table),
	}
}

func (c *Cqlx) Delete(table string) *DeleteBuilder {
	return &DeleteBuilder{
		cqlx: c,
		qb:   qb.Delete(table),
	}
}

func (c *Cqlx) InsertBuilderQuery(ib *TableInsertBuilder) *Queryx {
	stmt, names := ib.ToCql()

	return &Queryx{
		c.Session.Query(stmt, names),
	}
}

func (c *Cqlx) SelectBuilderQuery(sb *TableSelectBuilder) *Queryx {
	stmt, names := sb.ToCql()

	return &Queryx{
		c.Session.Query(stmt, names),
	}
}

func (c *Cqlx) UpdateBuilderQuery(ub *TableUpdateBuilder) *Queryx {
	stmt, names := ub.ToCql()

	return &Queryx{
		c.Session.Query(stmt, names),
	}
}
