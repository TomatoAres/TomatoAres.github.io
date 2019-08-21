package main

func getInt() int {
	return 1
}

/*
常量是一个简单值的标识符，在程序运行时，不会被修改的量。

1. 字面值（literal）常量，它们的值即为它们本身，是无法被改变的

下边飘红的全使用的变量，可能变化
2. 常量的值必须在编译期间确定。因此不能将函数的返回值赋给常量，因为函数调用发生在运行期。
const b = math.Sqrt(4)//not allowed

3. 这里可以理解为常量是无类型的，但是在需要类型的场合，编译器会根据常量的值和上下文将常量转换为相应的类型。

4.若需要声明类型，声明时加上类型
*/
func main() {
	ch := make(chan int, 1)
	ch <- 1
	s := "polaris"

	const (
		c1 = len(s)
		//c1 = "polaris" // 可行
		c2 = len("polaris")
		c3 = len([...]int{1, 2})
		c4 = len(&[...]int{3, 4})
		c5 = len([...]int{<-ch, 4})
		c52 = len([...]int{make(chan int),})
		c6 = len([...]int{getInt(), 4})
	)

}