package pack

import "log"

func init() {
	log.Println("myValue", myValue)

	myValue = 66
}

var myValue = 100

func Look() int {
	return myValue
}
