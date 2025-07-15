package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

const connNum = 1000
const callNum = 10
const url = "http://server/api/play/cash/ping"
const logFile = "./onion_doz_listen_2.csv"

func main() {

	log.SetFlags(0)
	w := bytes.NewBuffer([]byte{})
	log.SetOutput(w)

	buf := make(chan error, connNum*callNum)
	for i := 0; i < 10; i++ {
		now := time.Now()
		for k := 0; k < connNum; k++ { // 連線數
			client := newClient()
			for j := 0; j < callNum; j++ { // 呼叫次數
				id := k*connNum + j + 1
				go func(id int, client *http.Client) {
					buf <- curl(id, client)
				}(id, client)
			}
		}

		ok := 0
		for m := 0; m < connNum*callNum; m++ {
			err := <-buf
			if err != nil {
				ok++
			}
		}
		excursion := time.Since(now)
		tpr := excursion / (connNum * callNum)
		log.Printf("@%d,ConnNum:%d,CallNum:%d,Failed:%d,Excursion:%s,TPR:%s\n=====",
			i+1, connNum, connNum*callNum, ok, excursion, tpr,
		)
		time.Sleep(time.Second)
	}

	ioutil.WriteFile(logFile, w.Bytes(), 0777)
}

func curl(id int, client *http.Client) error {
	now := time.Now()
	res, err := http.Get(url + "?id=" + strconv.Itoa(id))
	if err != nil {
		return err
	}
	resTime := time.Now()

	if res.StatusCode != 200 {
		return fmt.Errorf("Status Code is %d", res.StatusCode)
	}

	st := res.Header.Get("t")
	reqTime, _ := time.Parse(time.RFC3339Nano, st)

	connExcursion := reqTime.Sub(now)
	reqExcursion := resTime.Sub(reqTime)
	f := "04:05.999999999"

	log.Printf(
		"#%d,%s,%s,%s,%s,%s",
		id,
		now.Format(f),
		reqTime.Format(f),
		resTime.Format(f),
		reqExcursion,
		connExcursion,
	)

	return nil
}

func newClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: func() func(network, addr string) (net.Conn, error) {
				var n net.Conn
				var e error
				return func(network, addr string) (net.Conn, error) {
					if n == nil {
						n, e = net.Dial(network, addr)
					}
					return n, e
				}
			}(),
		},
	}
}
