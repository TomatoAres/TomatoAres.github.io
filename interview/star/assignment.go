package main

import "fmt"
/*
多重赋值
先计算=左边，再进行赋值
*/
func main() {
	data, i := [4]int{0, 1, 2, 3}, 0
	i, data[i] = 1, 10
	fmt.Println(i, data)
}
