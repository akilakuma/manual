package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
	// "github.com/go-playground/validator"
	// "github.com/urfave/cli"
	// "github.com/joho/godotenv"
)

var conf IntergrationConfig

func main() {
	
	auto()

	iter()

}

func iter() {

	v := reflect.ValueOf(conf)

	var data []CCconfig

	for i := 0; i < v.NumField(); i++ {

		subV := v.Field(i).Interface()

		log.Println("value", subV)

		f := reflect.TypeOf(subV)

		if f.Kind() == reflect.Slice {

			s := reflect.Indirect(reflect.ValueOf(subV))

			for j := 0; j < s.Len(); j++ {
				//fmt.Println(s.Index(j))

				a := s.Index(j)

				fields := reflect.TypeOf(a.Interface())
				//fmt.Printf("%#v", a)
				fmt.Println()

				for k := 0; k < a.NumField(); k++ {
					field := fields.Field(k)
					log.Println("name:", field.Name, "value:", a.Field(k).Interface())
					temp := CCconfig{Key: field.Name, Value: a.Field(k).Interface()}
					data = append(data, temp)
				}
			}

		} else if f.Kind() == reflect.Struct {
			//log.Println("struct")

			fields := reflect.TypeOf(subV)
			values := reflect.ValueOf(subV)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				field := fields.Field(i)
				value := values.Field(i)
				fmt.Print("Type:", field.Type, ",", field.Name, "=", value, "\n")
			}
		} else {
			//log.Println("map")
		}
	}

	write(data)
}

func auto() {
	// 讀取設定檔案
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// viper.AutomaticEnv()
	// viper.SetEnvPrefix("TIGER")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//
	//var conf IntergrationConfig
	// // 根據mapstructure 取得參數
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Println(err)
	}

	//fmt.Printf("%#v", viper.GetStringMap("const"))

	//conf.Constant = viper.GetStringMap("const")

	fmt.Println()
	//
	//fmt.Printf("%#v", conf)

}

// IntergrationConfig 整合性的設定
type IntergrationConfig struct {
	MySQL    []ParserMysql
	CQL      []ParserCql
	Redis    []ParserRedis
	GRPC     []ParserGRPC
	NSQ      ParserNSQ
	Constant map[string]interface{}
}

// ParserMysql MySQL 解析
type ParserMysql struct {
	Name    string
	Host    string
	MaxIdle int
	MaxConn int
}

// ParserCql CQL 解析
type ParserCql struct {
	Name        string
	Host        string
	Port        string
	Keyspace    string
	PageSize    int
	Timeout     int
	Consistency string
}

// ParserRedis Redis 解析
type ParserRedis struct {
	Name    string
	Host    string
	MaxIdle int
	MaxConn int
	Number  int
}

// ParserNSQ NSQ 解析
type ParserNSQ struct {
	NsqdLookupHost         string
	NsqdLookupPort         string
	NsqdHost               string
	NsqdPort               string
	NsqdMaxInFlight        string
	NsqdHandlerConcurrency string
}

// ParserGRPC grpc 解析
type ParserGRPC struct {
	Name string
	Host string
}

type CCconfig struct {
	Key   string
	Value interface{}
}

func write(data []CCconfig) {
	myFile := ".env"
	fout, err := os.Create(myFile) //fout,	err := os.OpenFile(myFile,os.O_CREATE,0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fout.Write([]byte("env:\n"))
	for _, v := range data {
		//outstr := fmt.Sprintf("%s:%d\n", "Hello world", i)
		fout.WriteString("  - name: " + strings.ToUpper(v.Key) + "\n")
		fout.WriteString("    value: " + fmt.Sprintf("%v", v.Value) + "\n")
		//fout.Write([]byte("abcd\n"))
	}
	fout.Close()
}
