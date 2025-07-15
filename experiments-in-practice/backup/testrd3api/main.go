package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	req, err := http.NewRequest("GET", "http://172.17.178.55:80/api/user/3161880/commissionable?start_date=2019-04-02&end_date=2019-04-02&platform_id=23", nil)
	if err != nil {
		log.Println(err)
	}
	req.Host = "prod.cs.rd3"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	if err != nil {
		log.Println(err)
	}

}
