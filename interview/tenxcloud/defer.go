package main

import "fmt"

func calculate(index string,a int,b int)int  {
	ret := a+b
	fmt.Println(index,a,b,ret)
	return ret
}

func main() {
	a := 10
	b := 20
	defer calculate("1",a,calculate("1",a,b))

	a = 1
	defer calculate("2",a,calculate("1",a,b))

	b =0
	return
}