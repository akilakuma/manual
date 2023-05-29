package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func main() {
	fmt.Println("init ...")
	InitBigQuery()
	fmt.Println("query ...")
	uList := []int64{1182954603623215104, 1195890192970412032, 1198852282819538944, 1201782728914104320}
	BQSelectGAUserById(uList)
	//uList2 := []int64{1580237188403040256, 1580237098405867520}
	//BQSelectPCUserById(uList2)
}

func BQSelectGAUserById(uList []int64) {
	var (
		prefix = "events_intraday_"
		// 試算對應的時間(抓前一個小時)
		date  = time.Now().Add(-1 * time.Hour).Format("20060102")
		sList = transInt64SliceToString(uList)
	)

	query := "select user_id, count(*) from `" + projectID + "." + DatasetGAH5Id + "." + prefix + date + "` " +
		" where user_id in ( " + sList + " )" +
		" group by user_id"

	fmt.Println(query)

	q := clientOs.Query(query)
	it, runErr := q.Read(clientOsCtx)
	if runErr != nil {

		fmt.Println("BQSelectGAUserById query run  fail:" + runErr.Error())

	}

	// Iterate through the results.
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done { // from "google.golang.org/api/iterator"
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		valStr := fmt.Sprint(values[0])
		fmt.Println(valStr)
	}

	return
}

func BQSelectPCUserById(uList []int64) {
	var (
		prefix = "events_intraday_"
		// 試算對應的時間(抓前一個小時)
		date  = time.Now().Add(-1 * time.Hour).Format("20060102")
		sList = transInt64SliceToString(uList)
	)

	query := "select user_id, count(*) from `" + projectID + "." + DatasetGAPCId + "." + prefix + date + "` " +
		" where user_id in ( " + sList + " )" +
		" group by user_id"

	fmt.Println(query)

	q := clientOs.Query(query)
	it, runErr := q.Read(clientOsCtx)
	if runErr != nil {

		fmt.Println("BQSelectGAUserById query run  fail:" + runErr.Error())

	}

	// Iterate through the results.
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done { // from "google.golang.org/api/iterator"
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		valStr := fmt.Sprint(values[0])
		fmt.Println(valStr)
	}

	return
}

func transInt64SliceToString(iSlice []int64) string {
	var s []string
	for _, uId := range iSlice {
		s = append(s, `"`+strconv.FormatInt(uId, 10)+`"`)
	}
	uIdList := strings.Join(s, ",")
	return uIdList
}
