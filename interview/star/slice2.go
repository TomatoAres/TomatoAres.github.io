package main

import "fmt"

func main() {
	arr := []int{2, 3, 4}

	fmt.Printf("函数前：%p\n", arr)
	printSlice(arr)
}

func printSlice(arr []int) {
	fmt.Printf("函数中：%p\n", arr)
}

/*两次输出一样吗？为什么？*/

/*函数传参时，拷贝的不是结构体指针，是对整个结构体进行了拷贝，
但结构体内部数组指针还是指向同一片内存*/