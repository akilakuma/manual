package main

import (
	"fmt"
)

func main() {
	// test1()

	// var a T
	// var b T
	// fmt.Println(a.Equal(b))
	test5()
}

func test1() {
	t := Teacher{}
	t.ShowA()
}

type People struct{}
type a int

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func test2() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 6, 7}
	// s1 = append(s1, s2 )
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

type BetData struct {
	date   string
	Amount float64
	ratio  int
}

/*
func test3(userID, gameID int) (BetData, bool) {
	bet, isOK := getBetAPIData(userID, gameID)
	if isOK {
		return bet, isOK
	}

	return nil, isOK
}
*/

type T int

func (t T) Equal(u T) bool { return t == u }

/*

func GetUser(userList [][]string) {

	userOjectMap := make(map[int][]byte, 0)

	index := 0

	// X個一組的 user清單
	for _, userSlice := range userList {

		v := url.Values{}
		for _, userID := range userSlice {
			v.Add("users[]", userID)
		}
		v.Add("fields[]", "currency")
		v.Add("fields[]", "username")
		v.Add("fields[]", "all_parents")

		go func(v url.Values, subIndex int) {

			selfValue := v
			data, s, apiErr := DurianConn.Get("http://"+Conf.API.IP+"/api/users", selfValue)

			if s != 200 {
				log.Println(s)
			}

			if apiErr != nil {
				log.Println(apiErr.Error())
			}
			userOjectMap[subIndex] = data

		}(v, index)
		index++
	}
}
*/

type Student struct {
	name string
}

func test5() {
	m := map[string]Student{"people": {"liyuechun"}}
	fmt.Println(m)
	fmt.Println(m["people"])
}
