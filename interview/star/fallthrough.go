package main

import (
	"fmt"
)

func number() int {
	num := 2 * 5
	return num
}

func t1() {

	switch num := number(); {
	case num > 7:
		fmt.Printf("%d is more than 7\n", num)
		fallthrough
	case num > 20:
		fmt.Printf("%d is more than 20\n", num)
	}
}

func t2()  {
	v := 10
	switch v {
	case 10:
		fmt.Println(10)
		fallthrough
	case 20:
		fmt.Println(20)
		fallthrough
	case 30:
		fmt.Println(30)
		fallthrough
	default:
		fmt.Println("default")
	}
}

/*Go语言switch语句的case后默认带break，匹配成功后不会继续匹配其他case，跳出switch语句。使用fallthrough之后可以强制执行后面的那个case代码。*/

func main() {
	fmt.Println("t1:")
	t1()

	fmt.Println()

	fmt.Println("t2:")
	t2()
}