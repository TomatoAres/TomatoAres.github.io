package main

import "fmt"
type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	var stu student
	for _, stu = range stus {
		// 考点1
		//fmt.Println(stu)
		//fmt.Printf("%p\n",&stu)
		m[stu.Name] = &stu
	}

	for k,v := range m{
		// 考点2： key 细心程度 ，3 占位符
		fmt.Printf("key=%s,value=%v\n",k,v)
	}
}

func main() {
	parseStudent()
}

//out:
//key=zhou,value=&{wang 22}
//key=li,value=&{wang 22}
//key=wang,value=&{wang 22}
