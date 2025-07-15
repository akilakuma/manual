package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	dayForTotalToUse := time.Now().UTC().Add(-4 * time.Hour).Format("2006-01-02")

	convertTime, _ := time.Parse("2006-01-02", dayForTotalToUse)

	fmt.Println(convertTime)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
