package scylla

import (
	"pika.rdtech.vip/eden-lib/gocqlx/v2"
)

func Unsafe() {
	gocqlx.DefaultUnsafe = true
}

// Iterx is a wrapper around gocql.Iter which adds struct scanning capabilities.
type Iterx struct {
	*gocqlx.Iterx
}

// Unsafe forces the iterator to ignore missing fields. By default when scanning
// a struct if result row has a column that cannot be mapped to any destination
// field an error is reported. With unsafe such columns are ignored.
func (iter *Iterx) Unsafe() *Iterx {
	iter.Iterx.Unsafe()
	return iter
}

// StructOnly forces the iterator to treat a single-argument struct as
// non-scannable. This is is useful if you need to scan a row into a struct
// that also implements gocql.UDTUnmarshaler or in rare cases gocql.Unmarshaler.
func (iter *Iterx) StructOnly() *Iterx {
	iter.Iterx.StructOnly()

	return iter
}

// Get scans first row into a destination and closes the iterator.
//
// If the destination type is a struct pointer, then StructScan will be
// used.
// If the destination is some other type, then the row must only have one column
// which can scan into that type.
// This includes types that implement gocql.Unmarshaler and gocql.UDTUnmarshaler.
//
// If you'd like to treat a type that implements gocql.Unmarshaler or
// gocql.UDTUnmarshaler as an ordinary struct you should call
// StructOnly().Get(dest) instead.
//
// If no rows were selected, ErrNotFound is returned.
func (iter *Iterx) Get(dest interface{}) error {
	return iter.Iterx.Get(dest)
}

// Select scans all rows into a destination, which must be a pointer to slice
// of any type, and closes the iterator.
//
// If the destination slice type is a struct, then StructScan will be used
// on each row.
// If the destination is some other type, then each row must only have one
// column which can scan into that type.
// This includes types that implement gocql.Unmarshaler and gocql.UDTUnmarshaler.
//
// If you'd like to treat a type that implements gocql.Unmarshaler or
// gocql.UDTUnmarshaler as an ordinary struct you should call
// StructOnly().Select(dest) instead.
//
// If no rows were selected, ErrNotFound is NOT returned.
func (iter *Iterx) Select(dest interface{}) error {
	return iter.Iterx.Select(dest)
}

// StructScan is like gocql.Iter.Scan, but scans a single row into a single
// struct. Use this and iterate manually when the memory load of Select() might
// be prohibitive. StructScan caches the reflect work of matching up column
// positions to fields to avoid that overhead per scan, which means it is not
// safe to run StructScan on the same Iterx instance with different struct
// types.
func (iter *Iterx) StructScan(dest interface{}) bool {
	return iter.Iterx.StructScan(dest)
}

// Scan consumes the next row of the iterator and copies the columns of the
// current row into the values pointed at by dest. Use nil as a dest value
// to skip the corresponding column. Scan might send additional queries
// to the database to retrieve the next set of rows if paging was enabled.
//
// Scan returns true if the row was successfully unmarshaled or false if the
// end of the result set was reached or if an error occurred. Close should
// be called afterwards to retrieve any potential errors.
func (iter *Iterx) Scan(dest ...interface{}) bool {
	return iter.Iterx.Scan(dest)
}

// Close closes the iterator and returns any errors that happened during
// the query or the iteration.
func (iter *Iterx) Close() error {
	return iter.Iterx.Close()
}
