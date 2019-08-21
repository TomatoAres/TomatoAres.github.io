package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 类似python 类方法
func (*Person) Study(lang string) {
	fmt.Println("study language:", lang)
}

func (p Person) GetName() string  {
	return p.Name
}

func main() {
	var p *Person
	p.Study("Go")
	p.GetName()
}