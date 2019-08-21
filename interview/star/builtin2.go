package main

import "fmt"

/*x := make([]int, 0, 1)

那么：
1）append(x, 1)
2）x = append(x, 1) */

/*
(builtin)内置函数，如果有返回值，不允许不接收（_也是一种接收）。
*/
func t1()  {
	x := make([]int,0,1)
	append(x,1)
	fmt.Println(x)
}

func t2()  {
	x:= make([]int,0,1)
	x = append(x,1)
	fmt.Println(x)

}

func main() {
	t1()
	t2()
}