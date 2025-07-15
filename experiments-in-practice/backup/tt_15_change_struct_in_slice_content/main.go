package main

import "log"

type TargetS struct {
	A int
	B int
}

func main() {

	var targetS []TargetS

	for i := 0; i < 5; i++ {
		tmp := TargetS{
			A: 1,
		}
		targetS = append(targetS, tmp)
	}

	for i := 0; i < 5; i++ {
		for _, v := range targetS {
			if v.A == 3 {
				v.B = 3
			}
		}
	}

	log.Println(targetS)
}
