package scylla

import (
	"pika.rdtech.vip/eden-lib/gocqlx/v2/qb"
)

type SelectBuilder struct {
	qbs  *qb.SelectBuilder
	cqlx *Cqlx
}

const (
	// ASC is ascending order
	ASC qb.Order = true
	// DESC is descending order
	DESC = false
)

func (sb *SelectBuilder) ToCql() (stmt string, names []string) {
	return sb.qbs.ToCql()
}

func (sb *SelectBuilder) Query() *Queryx {
	stmt, names := sb.qbs.ToCql()

	return &Queryx{
		sb.cqlx.Session.Query(stmt, names),
	}
}

func (sb *SelectBuilder) Columns(columns ...string) *SelectBuilder {
	sb.qbs.Columns(columns...)

	return sb
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.qbs.From(table)

	return sb
}

func (sb *SelectBuilder) Json() *SelectBuilder {
	sb.qbs.Json()

	return sb
}

func As(column, name string) string {
	return column + " AS " + name
}

func (sb *SelectBuilder) Distinct(columns ...string) *SelectBuilder {
	sb.qbs.Distinct(columns...)

	return sb
}

func (sb *SelectBuilder) Where(w ...qb.Cmp) *SelectBuilder {
	sb.qbs.Where(w...)

	return sb
}

func (sb *SelectBuilder) GroupBy(columns ...string) *SelectBuilder {
	sb.qbs.GroupBy(columns...)

	return sb
}

func (sb *SelectBuilder) OrderBy(column string, o qb.Order) *SelectBuilder {
	sb.qbs.OrderBy(column, o)

	return sb
}

func (sb *SelectBuilder) Limit(limit uint) *SelectBuilder {
	sb.qbs.Limit(limit)

	return sb
}

func (sb *SelectBuilder) LimitPerPartition(limit uint) *SelectBuilder {
	sb.qbs.LimitPerPartition(limit)

	return sb
}

func (sb *SelectBuilder) AllowFiltering() *SelectBuilder {
	sb.qbs.AllowFiltering()

	return sb
}

// func (sb *SelectBuilder) BypassCache() *SelectBuilder {
// 	sb.qbs.BypassCache()

// 	return sb
// }

func (sb *SelectBuilder) Count(column string) *SelectBuilder {
	sb.qbs.Count(column)

	return sb
}

func (sb *SelectBuilder) CountAll() *SelectBuilder {
	sb.qbs.CountAll()

	return sb
}

func (sb *SelectBuilder) Min(column string) *SelectBuilder {
	sb.qbs.Min(column)

	return sb
}

func (sb *SelectBuilder) Max(column string) *SelectBuilder {
	sb.qbs.Max(column)

	return sb
}

func (sb *SelectBuilder) Avg(column string) *SelectBuilder {
	sb.qbs.Avg(column)

	return sb
}

func (sb *SelectBuilder) Sum(column string) *SelectBuilder {
	sb.qbs.Sum(column)

	return sb
}

func (sb *SelectBuilder) fn(name, column string) {
	sb.fn(name, column)
}
