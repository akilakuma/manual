package main

import (
	log "log"
	"time"
)

type Result struct {
	Machine string    `json:"machine"`
	Target  string    `json:"target"`
	Message string    `json:"message"`
	DoneAt  time.Time `json:"done_at"`
}

func main() {
	r := Result{
		Machine: "hello",
		Target:  "aaa.com",
		DoneAt:  time.Now(),
	}

	log.Println(r)
	newTime := time.Now()
	a := newTime.Add(1 * time.Hour)

	log.Println(a)

	if r.DoneAt.Before(a) {
		log.Println("before")
	} else {
		log.Println("after")
	}

}
