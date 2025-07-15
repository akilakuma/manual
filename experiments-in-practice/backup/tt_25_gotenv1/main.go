package main

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	env := os.Getenv("ENV")

	log.Println("the tt25 environment is:", env, " expect: alpha")

}
