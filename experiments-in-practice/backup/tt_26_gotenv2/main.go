package main

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	env := os.Getenv("ENV")

	log.Println("the tt26 environment is:", env, " expect: lembda")

}
