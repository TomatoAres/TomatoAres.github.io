package main

import "fmt"

/*
out:
0x5851d8
0x5851d8
true
TODO 暂时无法理解
*/
func main() {
	a := &struct {}{}
	b := &struct {}{}

	fmt.Printf("%p\n",a)
	fmt.Printf("%p\n",b)

	fmt.Println(a==b)
}
