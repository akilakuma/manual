package json

import (
	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
)

var js = jsoniter.ConfigCompatibleWithStandardLibrary

// ValueType Data types available in valid JSON data.
type ValueType int

const (
	NotExist = ValueType(iota)
	String
	Number
	Object
	Array
	Boolean
	Null
	Unknown
)

func (vt ValueType) String() string {
	switch vt {
	case NotExist:
		return "non-existent"
	case String:
		return "string"
	case Number:
		return "number"
	case Object:
		return "object"
	case Array:
		return "array"
	case Boolean:
		return "boolean"
	case Null:
		return "null"
	default:
		return "unknown"
	}
}

/*
ParserGet - Receives data structure, and key path to extract value from.

Returns:
`value` - Pointer to original data structure containing key value, or just empty slice if nothing found or error
`dataType` -    Can be: `NotExist`, `String`, `Number`, `Object`, `Array`, `Boolean` or `Null`
`offset` - Offset from provided data structure where key value ends. Used mostly internally, for example for `ArrayEach` helper.
`err` - If key not found or any other parsing issue it should return error. If key not found it also sets `dataType` to `NotExist`

Accept multiple keys to specify path to JSON value (in case of quering nested structures).
If no keys provided it will try to extract closest JSON value (simple ones or object/array), useful for reading streams or arrays, see `ArrayEach` implementation.
*/
func ParserGet(data []byte, keys ...string) (value []byte, dataType ValueType, offset int, err error) {
	v, d, o, e := jsonparser.Get(data, keys...)
	nd := getValueType(d)

	return v, nd, o, e
}

//GetString JsonParser GetString
//當key值字串中有[.]符號時請使用此func
//ex ip,mail..etc
func GetString(data []byte, keys ...string) (val string, err error) {
	val, err = jsonparser.GetString(data, keys...)
	return val, err
}

//GetBoolean JsonParser GetBoolean
//當key值字串中有[.]符號時請使用此func
//ex ip,mail..etc
func GetBoolean(data []byte, keys ...string) (val bool, err error) {
	val, err = jsonparser.GetBoolean(data, keys...)
	return val, err
}

//GetFloat JsonParser GetFloat
//當key值字串中有[.]符號時請使用此func
//ex ip,mail..etc
func GetFloat(data []byte, keys ...string) (val float64, err error) {
	val, err = jsonparser.GetFloat(data, keys...)
	return val, err
}

//GetInt JsonParser GetInt
//當key值字串中有[.]符號時請使用此func
//ex ip,mail..etc
func GetInt(data []byte, keys ...string) (val int64, err error) {
	val, err = jsonparser.GetInt(data, keys...)
	return val, err
}

//ObjectEach ObjectEach
func ObjectEach(data []byte, callback func(key []byte, value []byte, dataType ValueType, offset int) error, keys ...string) (err error) {
	return jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		newDt := getValueType(dataType)
		return callback(key, value, newDt, offset)
	}, keys...)
}

//ArrayEach ArrayEach
func ArrayEach(data []byte, cb func(value []byte, dataType ValueType, offset int, err error), keys ...string) {

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		newDt := getValueType(dataType)
		cb(value, newDt, offset, err)
	}, keys...)
}

func getValueType(in jsonparser.ValueType) (r ValueType) {
	switch in {
	case jsonparser.NotExist:
		r = NotExist
	case jsonparser.String:
		r = String
	case jsonparser.Number:
		r = Number
	case jsonparser.Object:
		r = Object
	case jsonparser.Array:
		r = Array
	case jsonparser.Boolean:
		r = Boolean
	case jsonparser.Null:
		r = Null
	case jsonparser.Unknown:
		r = Unknown
	default:
		r = Unknown
	}
	return

}

// ParseBoolean parses a Boolean ValueType into a Go bool (not particularly useful, but here for completeness)
func ParseBoolean(b []byte) (bool, error) {
	r, e := jsonparser.ParseBoolean(b)
	return r, e
}

// ParseString parses a String ValueType into a Go string (the main parsing work is unescaping the JSON string)
func ParseString(b []byte) (string, error) {
	r, e := jsonparser.ParseString(b)
	return r, e
}

// ParseFloat parses a Number ValueType into a Go float64
func ParseFloat(b []byte) (float64, error) {
	r, e := jsonparser.ParseFloat(b)
	return r, e
}

// ParseInt parses a Number ValueType into a Go int64
func ParseInt(b []byte) (int64, error) {
	r, e := jsonparser.ParseInt(b)
	return r, e
}

/*

Set - Receives existing data structure, path to set, and data to set at that key.

Returns:
`value` - modified byte array
`err` - On any parsing error

*/
func Set(data []byte, setValue []byte, keys ...string) (value []byte, err error) {
	value, err = jsonparser.Set(data, setValue, keys...)
	return
}

/*

Del - Receives existing data structure, path to delete.

Returns:
`data` - return modified data

*/
func Delete(data []byte, keys ...string) []byte {
	return jsonparser.Delete(data, keys...)
}

//EachKey EachKey
func EachKey(data []byte, cb func(i int, b []byte, v ValueType, e error), paths ...[]string) int {
	return jsonparser.EachKey(data, func(i int, b []byte, v jsonparser.ValueType, e error) {
		newDt := getValueType(v)
		cb(i, b, newDt, e)
	}, paths...)
}
