package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

// category 系列的id
var (
	clientOsCtx   context.Context
	clientOs      *bigquery.Client
	clientOsErr   error
	projectID     string
	DatasetGAH5Id string
	DatasetGAPCId string
)

func InitBigQuery() {
	projectID = "rd2-test"
	DatasetGAH5Id = "analytics_337368255"
	DatasetGAPCId = "analytics_339131859"

	clientOsCtx = context.Background()
	clientOs, clientOsErr = bigquery.NewClient(clientOsCtx, projectID)
	if clientOsErr != nil {
		fmt.Println("InitBigQuery bigquery client fail:" + clientOsErr.Error())
		return
	}
}
