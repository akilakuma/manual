package main

func main() {
	InitDB(EnvVariable{
		ExampleDBStr: "root:qwe123@tcp(127.0.0.1:3306)/example?charset=utf8&parseTime=True&loc=Local",
	})

	InsertRecord(6653, 10.5, "pika")

	GetTransRecord(6653, 0, "")
}
