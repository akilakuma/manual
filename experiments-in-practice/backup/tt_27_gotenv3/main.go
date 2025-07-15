package main

import (
	"log"
	"os"
)

func main() {

	env := os.Getenv("ENV")

	log.Println("the tt26 environment is:", env, " without 	gotenv.Load()")

}
