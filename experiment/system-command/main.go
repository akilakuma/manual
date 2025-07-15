package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	execPythonUpdateMA()
}

func execPythonUpdateMA() {

	cmd := exec.Command("curl", "http://localhost:8888/mem")
	cmd.Output()
	//time.Sleep(120 * time.Second)

	// 第一個參數是執行指令
	// 第二個以後是參數
	//cmd := exec.Command("/Users/shen_su/opt/anaconda3/envs/stock/bin/python", "/Users/shen_su/PycharmProjects/stock/update_ma.py")
	for i := 0; i < 10000; i++ {
		cmd := exec.Command("curl", "http://localhost:8888/ping")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(stdout))
		time.Sleep(1 * time.Second)
	}

	time.Sleep(10000 * time.Second)
}
