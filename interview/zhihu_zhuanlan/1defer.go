package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	panic("触发异常")

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
/*
打印后
打印中
打印前
panic: 触发异常——位置随机 TODO 怎么解释

解释：
return不是原子操作，往往是两步：
1. 赋值
2. 返回
defer调用会插到两者之间

这道题中，return 默认在panic后

*/