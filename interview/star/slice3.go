package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
//out：
//[0 0 0 0 0 1 2 3]