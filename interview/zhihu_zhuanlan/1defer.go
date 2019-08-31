package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
/*
打印后
打印中
打印前
panic: 触发异常——位置随机 

怎么解释?
Goland 编辑器的问题
1. Go 语言中的 标准输入、标准输出和标准错误 都是无缓冲的。这点和 C 语言不一样；
2. 当你遇到“诡异”问题时，通过终端方式验证下，很多时候可能编辑器，特别是 IDE 做了一些隐藏的事情。从上面的分析可以肯定，Goland 在处理标准输出和标准错误时做了额外的一些事情，具体是什么不得而知；
3. Go 中如果需要缓冲，请使用 bufio 包，特殊的需求，可以基于它进行扩展。

解释：
return不是原子操作，往往是两步：
1. 赋值
2. 返回
defer调用会插到两者之间

这道题中，return 默认在panic后

*/