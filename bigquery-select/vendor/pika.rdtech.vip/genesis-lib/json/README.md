# json
===

## 封裝套件
- https://github.com/tidwall/gjson
- https://github.com/json-iterator/go
- https://github.com/buger/jsonparser

## Method
* Valid(json string) bool
* ValidBytes(json []byte) bool
* GetMany(json string, path ...string) []Result
* GetManyBytes(json []byte, path ...string) []Result
* GetBytes(data []byte, path string) Resultstring) (r *Response)
* Parse(json string) Result
* ParseBytes(json []byte) Result
* Get(json string, path string) Result
* Marshal(v interface{}) ([]byte, error)
* Unmarshal(data []byte, v interface{}) error
* ParserGet(data []byte, keys ...string) (value []byte, dataType ValueType, offset int, err error)
* GetString(data []byte, keys ...string) (val string, err error)
* GetBoolean(data []byte, keys ...string) (val bool, err error)
* GetFloat(data []byte, keys ...string) (val float64, err error) 
* GetInt(data []byte, keys ...string) (val int64, err error)
* ObjectEach(data []byte, callback func(key []byte, value []byte, dataType ValueType, offset int) error, keys ...string) (err error)
* ArrayEach(data []byte, cb func(value []byte, dataType ValueType, offset int, err error), keys ...string)
* ParseBoolean(b []byte) (bool, error)
* ParseString(b []byte) (string, error)
* ParseFloat(b []byte) (float64, error)
* ParseInt(b []byte) (int64, error)
* Set(data []byte, setValue []byte, keys ...string) (value []byte, err error)
* Delete(data []byte, keys ...string) []byte
* EachKey(data []byte, cb func(i int, b []byte, v ValueType, e error), paths ...[]string) int

## Get/GetBytes
---
### 當key值有[.]這符號時，請使用ParserGet/GetString/GetBoolean/GetFloat/GetInt
### ex :{"127.0.0.1":"aaabbbcc"},{"imtest@gmail.com":"aabbccdd"}
### 因為Get/GetBytes是使用[.]符號進行分割字串
---

## func Get(json string, path string) Result
### 取得value，return值為Result strcut，查詢下層節點使用[.]進行串接，ex "persion.name.first"
```golang 
 get := json.Get(data, "person")
 value1 := get.String()
 value2 := get.Get("name").String()
 value3 := json.Get(data, "person.name.first").String() 
 value4 := json.Get(data, "person").String() 
 value5 := json.Get(data, "test1").Exists()
 value6 := json.Get(data, "test2").Array()
 value7 := json.Get(data, "test3").Float()
 value8 := json.Get(data, "test4").Bool()
 value9 := json.Get(data, "test5").Int()
```

## ParserGet(data []byte, keys ...string) (value []byte, dataType ValueType, offset int, err error)
### 取得value，當key值字串中有[.]符號時請使用此func
```golang 
a, b, c, d := json.ParserGet([]byte(data), "person", "name", "fullName")
```
## func GetString(data []byte, keys ...string) (val string, err error)
```golang 
 value1, error1 := json.GetString(data, "person", "github")
 value2, error2 := json.GetString(data, "127.0.0.1", "test")
```


## func ObjectEach(data []byte, callback func(key []byte, value []byte, dataType ValueType, offset int) error, keys ...string) (err error)   
```golang 
 json.ObjectEach(data, func(key []byte, value []byte, dataType json.ValueType, offset int) error {
		fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
		return nil
	}, "person", "name")
```

## func ArrayEach(data []byte, cb func(value []byte, dataType ValueType, offset int, err error), keys ...string) 
```golang 
json.ArrayEach(data, func(value []byte, dataType json.ValueType, offset int, err error) {
		fmt.Println(json.Get(value, "url"))
	}, "person", "avatars")
```


## func Valid(json string) bool  
```golang 
  exist := json.Valid(json)
```

## func ValidBytes(json []byte) bool  
```golang 
  exist := json.ValidBytes(json)
```

## GetMany(json string, path ...string) []Result 
```golang 
  arrResult := json.GetMany(data, "person.int1", "person.int2")
  for _, result := range arrResult {
	   
	}
```

## GetMany(json []byte, path ...string) []Result 
```golang 
  arrResult := json.GetMany([]byte(data), "person.int1", "person.int2")
  for _, result := range arrResult {
	   
	}
```