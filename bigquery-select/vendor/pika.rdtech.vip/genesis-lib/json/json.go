package json

import (
	j "encoding/json"
	"io"
	"time"

	"github.com/tidwall/gjson"
)

// Result represents a json value that is returned from Get().
type Result struct {
	js gjson.Result
}

//A Decoder reads and decodes JSON values from an input stream.
type Decoder struct {
	decoder *j.Decoder
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may
// read data from r beyond the JSON values requested.
func NewDecoder(r io.Reader) *Decoder {
	decoder := j.NewDecoder(r)
	return &Decoder{decoder: decoder}
}

// Decode reads the next JSON-encoded value from its
// input and stores it in the value pointed to by v.
//
// See the documentation for Unmarshal for details about
// the conversion of JSON into a Go value.
func (dec Decoder) Decode(v interface{}) error {
	return dec.decoder.Decode(v)
}

// String String
func (r Result) String() string {
	return r.js.String()
}

// Exists Exists
func (r Result) Exists() bool {
	return r.js.Exists()
}

// Array Array
func (r Result) Array() []Result {
	result := []Result{}
	gres := r.js.Array()
	for _, v := range gres {
		result = append(result, Result{
			js: v,
		})
	}

	return result
}

// Float returns an float64 representation.
func (r Result) Float() float64 {
	return r.js.Float()
}

// Bool returns an boolean representation.
func (r Result) Bool() bool {
	return r.js.Bool()
}

// Int returns an integer representation.
func (r Result) Int() int64 {
	return r.js.Int()
}

// Uint returns an unsigned integer representation.
func (r Result) Uint() uint64 {
	return r.js.Uint()
}

// Time returns a time.Time representation.
func (r Result) Time() time.Time {
	return r.js.Time()
}

// IsObject returns true if the result value is a JSON object.
func (r Result) IsObject() bool {
	return r.js.IsObject()
}

// IsArray returns true if the result value is a JSON array.
func (r Result) IsArray() bool {
	return r.js.IsArray()
}

// Map returns back an map of values. The result should be a JSON array.
func (r Result) Map() map[string]Result {
	ret := make(map[string]Result)
	mp := r.js.Map()
	for k, v := range mp {
		mpv := Result{js: v}
		ret[k] = mpv
	}
	return ret
}

// Get searches result for the specified path.
// The result should be a JSON array or object.
// 註：當key值有包含[.]符號時，ex ip,mail..etc，會被判斷成下層節點，請改用GetString/GetInt/GetBoolean/GetFloat等func
func (r Result) Get(path string) Result {
	gres := gjson.Get(r.js.Raw, path)
	return Result{js: gres}
}

// Value returns one of these types:
//
//	bool, for JSON booleans
//	float64, for JSON numbers
//	Number, for JSON numbers
//	string, for JSON string literals
//	nil, for JSON null
//	map[string]interface{}, for JSON objects
//	[]interface{}, for JSON arrays
//
func (r Result) Value() interface{} {
	return r.js.Value()
}

//Valid Valid
func Valid(json string) bool {
	return gjson.Valid(json)
}

//ValidBytes ValidBytes
func ValidBytes(json []byte) bool {
	return gjson.ValidBytes(json)
}

//GetMany GetMany
// 註：當key值有包含[.]符號時，ex ip,mail..etc，會被判斷成下層節點，請改用GetString/GetInt/GetBoolean/GetFloat等func
func GetMany(json string, path ...string) []Result {
	ret := []Result{}
	gres := gjson.GetMany(json, path...)
	for _, v := range gres {
		gs := Result{
			js: v,
		}
		ret = append(ret, gs)
	}
	return ret
}

//GetManyBytes GetManyBytes
// 註：當key值有包含[.]符號時，ex ip,mail..etc，會被判斷成下層節點，請改用GetString/GetInt/GetBoolean/GetFloat等func
func GetManyBytes(json []byte, path ...string) []Result {
	ret := []Result{}
	gres := gjson.GetManyBytes(json, path...)
	for _, v := range gres {
		gs := Result{
			js: v,
		}
		ret = append(ret, gs)
	}
	return ret
}

// ForEach ForEach
func (r Result) ForEach(fc func(upkey, upvalue Result) bool) {
	r.js.ForEach(func(key, value gjson.Result) bool {
		inkey := Result{js: key}
		inval := Result{js: value}

		return fc(inkey, inval)
	})
}

// GetBytes GetBytes
func GetBytes(data []byte, path string) Result {
	res := gjson.GetBytes(data, path)

	return Result{
		js: res,
	}
}

//Parse Parse
func Parse(json string) Result {
	gres := gjson.Parse(json)
	return Result{js: gres}
}

//ParseBytes ParseBytes
func ParseBytes(json []byte) Result {
	gres := gjson.ParseBytes(json)
	return Result{js: gres}
}

//Get Get
func Get(json string, path string) Result {
	gres := gjson.Get(json, path)
	return Result{js: gres}
}

//--------------json-iterator-------------------

//Marshal Marshal
func Marshal(v interface{}) ([]byte, error) {
	ret, err := js.Marshal(&v)
	return ret, err
}

//MarshalToString MarshalToString
// func MarshalToString(v interface{}) (string, error) {
// 	ret, err := js.MarshalToString(&v)
// 	return ret, err
// }

//UnmarshalFromString UnmarshalFromString
// func UnmarshalFromString(str string, v interface{}) error {
// 	return js.UnmarshalFromString(str, &v)
// }

//Unmarshal Unmarshal
func Unmarshal(data []byte, v interface{}) error {
	return js.Unmarshal(data, &v)
}
