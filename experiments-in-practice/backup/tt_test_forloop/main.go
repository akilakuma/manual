package main

import (
	"log"
)

func main() {

	var i int
	for {
		println(i)

		i++
		select {
		default:
			if i == 10 {
				return
			}
		}

		log.Println("yo")
	}

}

type ReadCloser interface {
	Read(b []byte) (n int, err error)
	Close()
}

func ReadAndClose(r ReadCloser, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}
