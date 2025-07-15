package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func LoadCsv(filePath string) [][]string {

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777) // os.O_RDONLY 表示只讀、0777 表示(owner/group/other)權限
	if err != nil {
		log.Fatalln("找不到CSV檔案路徑:", filePath, err)
	}

	// read
	r := csv.NewReader(file)
	r.Comma = ',' // 以何種字元作分隔，預設為`,`。所以這裡可拿掉這行
	var data [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		data = append(data, record)
	}
	return data
}

func GenerateFileByNameAndCol(name string, dataPara, col [][]string) {

	var writer *csv.Writer

	// 打開檔案，沒有的話建一個
	file, err := os.OpenFile(name+".csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create(name + ".csv")
		checkError("Cannot create file", err)

		writer = csv.NewWriter(file)
		for _, value := range col {
			err := writer.Write(value)
			checkError("Cannot write to "+name+" file", err)
		}
	} else {
		writer = csv.NewWriter(file)
	}

	defer file.Close()
	defer writer.Flush()

	for _, value := range dataPara {
		err := writer.Write(value)
		checkError("Cannot write to result file", err)
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
