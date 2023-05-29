package scylla

import "pika.rdtech.vip/eden-lib/gocqlx/v2/qb"

// Metadata represents table schema.
type Metadata struct {
	Name    string
	Columns []string
	PartKey []string
	SortKey []string
}

type cql struct {
	stmt  string
	names []string
}

// Table allows for simple CRUD operations, it's backed by query builders from
// gocqlx/qb package.
type Table struct {
	metadata      Metadata
	primaryKeyCmp []qb.Cmp
	partKeyCmp    []qb.Cmp

	get    cql
	sel    cql
	insert cql
}

// New creates new Table based on table schema read from Metadata.
func NewTable(m Metadata) *Table { // nolint: gocritic
	t := &Table{
		metadata: m,
	}

	// prepare primary and partition key comparators
	t.primaryKeyCmp = make([]qb.Cmp, 0, len(m.PartKey)+len(m.SortKey))
	for _, k := range m.PartKey {
		t.primaryKeyCmp = append(t.primaryKeyCmp, qb.Eq(k))
	}

	for _, k := range m.SortKey {
		t.primaryKeyCmp = append(t.primaryKeyCmp, qb.Eq(k))
	}

	t.partKeyCmp = make([]qb.Cmp, len(m.PartKey))
	copy(t.partKeyCmp, t.primaryKeyCmp[:len(t.metadata.PartKey)])

	// prepare get stmt
	t.get.stmt, t.get.names = qb.Select(m.Name).Where(t.primaryKeyCmp...).ToCql()
	// prepare select stmt
	t.sel.stmt, t.sel.names = qb.Select(m.Name).Where(t.partKeyCmp...).ToCql()
	// prepare insert stmt
	t.insert.stmt, t.insert.names = qb.Insert(m.Name).Columns(m.Columns...).ToCql()

	return t
}

// Metadata returns copy of table metadata.
func (t *Table) Metadata() Metadata {
	return t.metadata
}

// PrimaryKeyCmp returns copy of table's primaryKeyCmp.
func (t *Table) PrimaryKeyCmp() []qb.Cmp {
	primaryKeyCmp := make([]qb.Cmp, len(t.primaryKeyCmp))
	copy(primaryKeyCmp, t.primaryKeyCmp)

	return primaryKeyCmp
}

// Name returns table name.
func (t *Table) Name() string {
	return t.metadata.Name
}

// SelectBuilder returns a builder initialised to select by partition key
// statement.
func (t *Table) SelectBuilder(columns ...string) *TableSelectBuilder {
	return &TableSelectBuilder{
		qb.Select(t.metadata.Name).Columns(columns...).Where(t.partKeyCmp...),
	}
}

// InsertBuilder
func (t *Table) InsertBuilder() *TableInsertBuilder {
	return &TableInsertBuilder{
		qb.Insert(t.metadata.Name).Columns(t.metadata.Columns...),
	}
}

// UpdateBuilder returns a builder initialised to update by primary key statement.
func (t *Table) UpdateBuilder(columns ...string) *TableUpdateBuilder {
	return &TableUpdateBuilder{
		qb.Update(t.metadata.Name).Set(columns...).Where(t.primaryKeyCmp...),
	}
}

// DeleteBuilder returns a builder initialised to delete by primary key statement.
func (t *Table) DeleteBuilder(columns ...string) *TableDeleteBuilder {
	return &TableDeleteBuilder{
		qb.Delete(t.metadata.Name).Columns(columns...).Where(t.primaryKeyCmp...),
	}
}

type TableSelectBuilder struct {
	*qb.SelectBuilder
}

func (tsb *TableSelectBuilder) ToCql() (stmt string, names []string) {
	return tsb.SelectBuilder.ToCql()
}

type TableUpdateBuilder struct {
	*qb.UpdateBuilder
}

func (tub *TableUpdateBuilder) ToCql() (stmt string, names []string) {
	return tub.UpdateBuilder.ToCql()
}

type TableDeleteBuilder struct {
	*qb.DeleteBuilder
}

func (tdb *TableDeleteBuilder) ToCql() (stmt string, names []string) {
	return tdb.DeleteBuilder.ToCql()
}
