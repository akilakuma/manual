package main

import (
	"fmt"
	"os/exec"
)

func main() {
	execPythonUpdateMA()
}

func execPythonUpdateMA() {
	// 第一個參數是執行指令
	// 第二個以後是參數
	cmd := exec.Command("/Users/shen_su/opt/anaconda3/envs/stock/bin/python", "/Users/shen_su/PycharmProjects/stock/update_ma.py")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
