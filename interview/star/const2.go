//以下代码是否有问题？为什么？
package main

import (
	"fmt"
)

func main() {
	const a = 2
	const b int32 = 3

	var c int = a
	var d int = b
	fmt.Println(c, d)
}
