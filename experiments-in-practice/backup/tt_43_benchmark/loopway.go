package main

var (
	abc = []string{"9", "2", "4", "1"}
	t   = "6"
	B   int
)

func loopway() {

	for i := 0; i < 4; i++ {
		if abc[i] == t {
			B++
		}
	}
}
