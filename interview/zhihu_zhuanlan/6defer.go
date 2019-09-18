package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*类似闭包

defer 进栈之前 执行参数
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4

我认为：
20 0 2 2
2 0 2 2
10 1 2 3
1 1 3 4
*/