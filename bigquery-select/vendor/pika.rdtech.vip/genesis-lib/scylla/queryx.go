package scylla

import (
	"context"

	"pika.rdtech.vip/eden-lib/gocqlx/v2"
	"github.com/syhlion/gocql"
)

type Queryx struct {
	*gocqlx.Queryx
}

// BindStruct binds query named parameters to values from arg using mapper. If
// value cannot be found error is reported.
func (q *Queryx) BindStruct(arg interface{}) *Queryx {
	q.Queryx.BindStruct(arg)

	return q
}

// BindStructMap binds query named parameters to values from arg0 and arg1
// using a mapper. If value cannot be found in arg0 it's looked up in arg1
// before reporting an error.
func (q *Queryx) BindStructMap(arg0 interface{}, arg1 map[string]interface{}) *Queryx {
	q.Queryx.BindStructMap(arg0, arg1)

	return q
}

// BindMap binds query named parameters using map.
func (q *Queryx) BindMap(arg map[string]interface{}) *Queryx {
	q.Queryx.BindMap(arg)

	return q
}

// Bind sets query arguments of query. This can also be used to rebind new query arguments
// to an existing query instance.
func (q *Queryx) Bind(v ...interface{}) *Queryx {
	q.Queryx.Bind(v...)

	return q
}

// Err returns any binding errors.
func (q *Queryx) Err() error {
	return q.Queryx.Err()
}

// Exec executes the query without returning any rows.
func (q *Queryx) Exec() error {
	return q.Queryx.Exec()
}

// ExecRelease calls Exec and releases the query, a released query cannot be
// reused.
func (q *Queryx) ExecRelease() error {
	return q.Queryx.ExecRelease()
}

// ExecCAS executes the Lightweight Transaction query, returns whether query was applied.
// See: https://docs.scylladb.com/using-scylla/lwt/ for more details.
func (q *Queryx) ExecCAS() (applied bool, err error) {
	return q.Queryx.ExecCAS()
}

// ExecCASRelease calls ExecCAS and releases the query, a released query cannot be
// reused.
func (q *Queryx) ExecCASRelease() (bool, error) {
	return q.Queryx.ExecCASRelease()
}

// Get scans first row into a destination and closes the iterator.
//
// If the destination type is a struct pointer, then Iter.StructScan will be
// used.
// If the destination is some other type, then the row must only have one column
// which can scan into that type.
// This includes types that implement gocql.Unmarshaler and gocql.UDTUnmarshaler.
//
// If you'd like to treat a type that implements gocql.Unmarshaler or
// gocql.UDTUnmarshaler as an ordinary struct you should call
// Iter().StructOnly().Get(dest) instead.
//
// If no rows were selected, ErrNotFound is returned.
func (q *Queryx) Get(dest interface{}) error {
	return q.Queryx.Get(dest)
}

// GetRelease calls Get and releases the query, a released query cannot be
// reused.
func (q *Queryx) GetRelease(dest interface{}) error {
	return q.Queryx.GetRelease(dest)
}

// GetCAS executes a lightweight transaction.
// If the transaction fails because the existing values did not match,
// the previous values will be stored in dest object.
// See: https://docs.scylladb.com/using-scylla/lwt/ for more details.
func (q *Queryx) GetCAS(dest interface{}) (applied bool, err error) {
	return q.Queryx.GetCAS(dest)
}

// GetCASRelease calls GetCAS and releases the query, a released query cannot be
// reused.
func (q *Queryx) GetCASRelease(dest interface{}) (bool, error) {
	return q.Queryx.GetCASRelease(dest)
}

// Select scans all rows into a destination, which must be a pointer to slice
// of any type, and closes the iterator.
//
// If the destination slice type is a struct, then Iter.StructScan will be used
// on each row.
// If the destination is some other type, then each row must only have one
// column which can scan into that type.
// This includes types that implement gocql.Unmarshaler and gocql.UDTUnmarshaler.
//
// If you'd like to treat a type that implements gocql.Unmarshaler or
// gocql.UDTUnmarshaler as an ordinary struct you should call
// Iter().StructOnly().Select(dest) instead.
//
// If no rows were selected, ErrNotFound is NOT returned.
func (q *Queryx) Select(dest interface{}) error {
	return q.Queryx.Select(dest)
}

// SelectRelease calls Select and releases the query, a released query cannot be
// reused.
func (q *Queryx) SelectRelease(dest interface{}) error {
	return q.Queryx.SelectRelease(dest)
}

// Iter returns Iterx instance for the query. It should be used when data is too
// big to be loaded with Select in order to do row by row iteration.
// See Iterx StructScan function.
func (q *Queryx) Iter() *Iterx {
	return &Iterx{
		q.Queryx.Iter(),
	}
}

// This file contains wrappers around gocql.Queryx that make Queryx expose the
// same interface but return *Queryx, this should be inlined by compiler.
//----------------------------------------------------------------------------

// Consistency sets the consistency level for this query. If no consistency
// level have been set, the default consistency level of the cluster
// is used.
func (q *Queryx) Consistency(c gocql.Consistency) *Queryx {
	q.Queryx.Consistency(c)

	return q
}

// PageSize will tell the iterator to fetch the result in pages of size n.
// This is useful for iterating over large result sets, but setting the
// page size too low might decrease the performance. This feature is only
// available in Cassandra 2 and onwards.
func (q *Queryx) PageSize(n int) *Queryx {
	q.Queryx.PageSize(n)

	return q
}

// WithTimestamp will enable the with default timestamp flag on the query
// like DefaultTimestamp does. But also allows to define value for timestamp.
// It works the same way as USING TIMESTAMP in the query itself, but
// should not break prepared query optimization
//
// Only available on protocol >= 3
func (q *Queryx) WithTimestamp(timestamp int64) *Queryx {
	q.Queryx.WithTimestamp(timestamp)

	return q
}

// WithContext returns a shallow copy of q with its context
// set to ctx.
//
// The provided context controls the entire lifetime of executing a
// query, queries will be canceled and return once the context is
// canceled.
func (q *Queryx) WithContext(ctx context.Context) *Queryx {
	q.Queryx.WithContext(ctx)

	return q
}

// RetryPolicy sets the policy to use when retrying the query.
func (q *Queryx) RetryPolicy(r gocql.RetryPolicy) *Queryx {
	q.Queryx.RetryPolicy(r)

	return q
}
