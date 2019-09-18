package main

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

//2.
func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA()) // 0

	fmt.Println(increaseB()) //1
}
