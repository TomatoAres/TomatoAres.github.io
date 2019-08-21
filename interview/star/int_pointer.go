package main

import (
"fmt"
)

func t1() {
	var i int = 2
	p := &i

	*p,b :=  3,4

	//b := 4
	//q := &b

	//*p ,*q = 5,6
	//*p, b := 3, 4
	/*
	 会报错non-name *p on left side of :=
	:= 左边的变量只能是个简单变量，
	我的理解：*p 变量 声明过了，这里 :=  重新声明b，将*p重新声明不合适
	*/

	fmt.Println(*p, b)
}

func GetName() (string, error) {
	return "jack", nil
}
type  Person struct {
	Name string
}

func similar()  {
	var jk Person
	//jk.Name, err := GetName() // 不能使用简短声明来设置字段的值
	_,err := GetName()
	//jk.Name 是另一个类型中的类型，会报错
	fmt.Println(jk.Name,err)
}

func main() {
	t1()

	fmt.Println()

	similar()
}
