package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func PrintLocalDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 1 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}

	fmt.Println("connect done, use", conn.LocalAddr().String())

	return conn, err
}

func doGet(client *http.Client, url string, id int) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.ReadAll(resp.Body)
	log.Println(id)

	resp.Body.Close()

	// fmt.Printf("%d: %s -- %v\n", id, string(buf), err)
	// if err2 := resp.Body.Close(); err2 != nil {
	// 	fmt.Println(err2)
	// }
}

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			Dial:              PrintLocalDial,
			DisableKeepAlives: false,
		},
	}

	URL := "https://serholiu.com/go-http-client-keepalive"

	for {
		go doGet(client, URL, 1)
		go doGet(client, URL, 2)
		time.Sleep(5 * time.Second)
		log.Println()
	}
}
