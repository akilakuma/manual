package scylla

import (
	"pika.rdtech.vip/eden-lib/gocqlx/v2/qb"
)

type QbCmp struct {
	Cmp qb.Cmp
}

func Eq(column string) qb.Cmp {
	return qb.Eq(column)
}

func EqTuple(column string, count int) qb.Cmp {
	return qb.EqTuple(column, count)
}

func EqNamed(column, name string) qb.Cmp {
	return qb.EqNamed(column, name)
}

func EqLit(column, literal string) qb.Cmp {
	return qb.EqLit(column, literal)
}

// func EqFunc(column string, fn *Func) qb.Cmp {
// 	return qb.EqFunc(column, fn)
// }

func Ne(column string) qb.Cmp {
	return qb.Ne(column)
}

func NeTuple(column string, count int) qb.Cmp {
	return qb.NeTuple(column, count)
}

func NeNamed(column, name string) qb.Cmp {
	return qb.NeNamed(column, name)
}

func NeLit(column, literal string) qb.Cmp {
	return qb.NeLit(column, literal)
}

// func NeFunc(column string, fn *Func) qb.Cmp {
// 	return qb.NeFunc(column, fn)
// }

// NeFunc produces column!=someFunc(?...).
// func NeFunc(column string, fn *Func) qb.Cmp {
// 	return qb.(column, literal)
// }

// Lt produces column<?.
func Lt(column string) qb.Cmp {
	return qb.Lt(column)
}

// LtTuple produces column<(?,?,...) with count placeholders.
func LtTuple(column string, count int) qb.Cmp {
	return qb.LtTuple(column, count)
}

// LtNamed produces column<? with a custom parameter name.
func LtNamed(column, name string) qb.Cmp {
	return qb.LtNamed(column, name)
}

// LtLit produces column<literal and does not add a parameter to the query.
func LtLit(column, literal string) qb.Cmp {
	return qb.LtLit(column, literal)
}

// LtFunc produces column<someFunc(?...).
// func LtFunc(column string, fn *Func) qb.Cmp {
// 	return qb.(column, literal)
// }

// LtOrEq produces column<=?.
func LtOrEq(column string) qb.Cmp {
	return qb.LtOrEq(column)
}

// LtOrEqTuple produces column<=(?,?,...) with count placeholders.
func LtOrEqTuple(column string, count int) qb.Cmp {
	return qb.LtOrEqTuple(column, count)
}

// LtOrEqNamed produces column<=? with a custom parameter name.
func LtOrEqNamed(column, name string) qb.Cmp {
	return qb.LtOrEqNamed(column, name)
}

// LtOrEqLit produces column<=literal and does not add a parameter to the query.
func LtOrEqLit(column, literal string) qb.Cmp {
	return qb.LtOrEqLit(column, literal)
}

// LtOrEqFunc produces column<=someFunc(?...).
// func LtOrEqFunc(column string, fn *Func) qb.Cmp {
// 	return qb.LtOrEqFunc(column, literal)
// }

// Gt produces column>?.
func Gt(column string) qb.Cmp {
	return qb.Gt(column)
}

// GtTuple produces column>(?,?,...) with count placeholders.
func GtTuple(column string, count int) qb.Cmp {
	return qb.GtTuple(column, count)
}

// GtNamed produces column>? with a custom parameter name.
func GtNamed(column, name string) qb.Cmp {
	return qb.GtNamed(column, name)
}

// GtLit produces column>literal and does not add a parameter to the query.
func GtLit(column, literal string) qb.Cmp {
	return qb.GtLit(column, literal)
}

// GtFunc produces column>someFunc(?...).
// func GtFunc(column string, fn *Func) qb.Cmp {
// 	return qb.(column, literal)
// }

// GtOrEq produces column>=?.
func GtOrEq(column string) qb.Cmp {
	return qb.GtOrEq(column)
}

// GtOrEqTuple produces column>=(?,?,...) with count placeholders.
func GtOrEqTuple(column string, count int) qb.Cmp {
	return qb.GtOrEqTuple(column, count)
}

// GtOrEqNamed produces column>=? with a custom parameter name.
func GtOrEqNamed(column, name string) qb.Cmp {
	return qb.GtOrEqNamed(column, name)
}

// GtOrEqLit produces column>=literal and does not add a parameter to the query.
func GtOrEqLit(column, literal string) qb.Cmp {
	return qb.GtOrEqLit(column, literal)
}

// GtOrEqFunc produces column>=someFunc(?...).
// func GtOrEqFunc(column string, fn *Func) qb.Cmp {
// 	return qb.(column, literal)
// }

// In produces column IN ?.
func In(column string) qb.Cmp {
	return qb.In(column)
}

// InTuple produces column IN ?.
func InTuple(column string, count int) qb.Cmp {
	return qb.InTuple(column, count)
}

// InNamed produces column IN ? with a custom parameter name.
func InNamed(column, name string) qb.Cmp {
	return qb.InNamed(column, name)
}

// InLit produces column IN literal and does not add a parameter to the query.
func InLit(column, literal string) qb.Cmp {
	return qb.InLit(column, literal)
}

// Contains produces column CONTAINS ?.
func Contains(column string) qb.Cmp {
	return qb.Contains(column)
}

// ContainsTuple produces column CONTAINS (?,?,...) with count placeholders.
func ContainsTuple(column string, count int) qb.Cmp {
	return qb.ContainsTuple(column, count)
}

// ContainsKey produces column CONTAINS KEY ?.
func ContainsKey(column string) qb.Cmp {
	return qb.ContainsKey(column)
}

// ContainsKeyTuple produces column CONTAINS KEY (?,?,...) with count placehplders.
func ContainsKeyTuple(column string, count int) qb.Cmp {
	return qb.ContainsKeyTuple(column, count)
}

// ContainsNamed produces column CONTAINS ? with a custom parameter name.
func ContainsNamed(column, name string) qb.Cmp {
	return qb.ContainsNamed(column, name)
}

// ContainsKeyNamed produces column CONTAINS KEY ? with a custom parameter name.
func ContainsKeyNamed(column, name string) qb.Cmp {
	return qb.ContainsKeyNamed(column, name)
}

// ContainsLit produces column CONTAINS literal and does not add a parameter to the query.
func ContainsLit(column, literal string) qb.Cmp {
	return qb.ContainsLit(column, literal)
}

// Like produces column LIKE ?.
func Like(column string) qb.Cmp {
	return qb.Like(column)
}

// LikeTuple produces column LIKE (?,?,...) with count placeholders.
func LikeTuple(column string, count int) qb.Cmp {
	return qb.LikeTuple(column, count)
}
