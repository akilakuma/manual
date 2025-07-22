package main

/*
 sync.Pool 還回去之前應該先reset
*/

func main() {
	useBadPool()

	useGoodPool()
}
