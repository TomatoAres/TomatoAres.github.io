package main

func main() {
	var a [1]int
	c := a[:]
	println(c)
}

/*
需要看文章，进行反编译
 */
/*
package main

import "fmt"

func main() {
   var a [1]int
   c := a[:]
   fmt.Println(c)
}
 */