package main

import (
	"fmt"
	"log"
)

type UserInfo struct {
	A int
	B string
	C int
}

type UserInfoMap map[int]*UserInfo

func (m UserInfoMap) SetA(i, v int) {
	m[i].A = v
}

func (m *UserInfoMap) SetC(i, v int) {
	userMap := *m
	userMap[i].C = v
}

func main() {
	UserInfoList := make(map[int]*UserInfo, 0)

	UserInfoList[12345] = &UserInfo{10, "i am A", 100}
	UserInfoList[666] = &UserInfo{60, "i am B", 66}
	UserInfoMap(UserInfoList).SetA(666, 9)
	fmt.Println(UserInfoList[666])

	userList := make(UserInfoMap, 0)
	userList[777] = &UserInfo{10, "i am A", 100} // 沒有這行initial，SetC會死翹翹
	userList.SetC(777, 999)
	log.Println(userList[777])
}
