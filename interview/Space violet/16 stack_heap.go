package main

import "runtime/debug"

func main() {
	//slice := make([]string, 2, 4)
	map_ := make(map[int]int)
	Example(map_, "hello", 10)
}
func Example(m map[int]int, str string, i int) {
	debug.PrintStack()
}