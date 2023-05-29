package main

import "fmt"

func main() {
	InitDB(EnvVariable{
		OrderDBStr: "root:qwe123@tcp(127.0.0.1:3306)/order_statistic?charset=utf8&parseTime=True&loc=Local",
	})
	fmt.Println("db init")

	r := QueryFirst(1158982199926583296)
	fmt.Println(r)

}
