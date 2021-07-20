package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var FilePath = "draws.csv"


func main() {

	// 從csv讀取
	importCSV()
}

func importCSV() {
	file, err := os.OpenFile(FilePath, os.O_RDONLY, 0777) // os.O_RDONLY 表示只讀、0777 表示(owner/group/other)權限
	if err != nil {
		log.Fatalln("找不到CSV檔案路徑:", FilePath, err)
	}


	// read
	r := csv.NewReader(file)
	r.Comma = ',' // 以何種字元作分隔，預設為`,`。所以這裡可拿掉這行
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println( record)
	}
}

