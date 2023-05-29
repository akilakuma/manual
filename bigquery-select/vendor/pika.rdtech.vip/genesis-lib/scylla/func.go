package scylla

// Functions reference:
// https://cassandra.apache.org/doc/latest/cql/functions.html

// Func is a custom database function invocation that can be use in a comparator
// or update statement.
type Func struct {
	// function name
	Name string
	// name of the function parameters
	ParamNames []string
}

// Fn creates Func.
func Fn(name string, paramNames ...string) *Func {
	return &Func{
		Name:       name,
		ParamNames: paramNames,
	}
}

// MinTimeuuid produces minTimeuuid(?).
func MinTimeuuid(name string) *Func {
	return Fn("minTimeuuid", name)
}

// MaxTimeuuid produces maxTimeuuid(?).
func MaxTimeuuid(name string) *Func {
	return Fn("maxTimeuuid", name)
}

// Now produces now().
func Now() *Func {
	return Fn("now")
}
