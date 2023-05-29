package main

import (
	"fmt"
	"sort"
)

func main() {
	dataList := []userData{
		userData{123, 10},
		userData{111, 5},
		userData{222, 1},
	}
	sort.Sort(userDataSlice(dataList))
	for _, v := range dataList {
		fmt.Println(v.drawId, v.betGold)
	}

}

type userData struct {
	drawId  int64
	betGold float64
}

type userDataSlice []userData

func (e userDataSlice) Len() int {
	return len(e)
}

func (e userDataSlice) Less(i, j int) bool {
	return e[i].drawId < e[j].drawId
}

func (e userDataSlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
