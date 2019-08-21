package main

type Person struct {
	Age int
}

func (p *Person) GrowUp() {
	p.Age++
}

func main() {
	m := map[string]Person{
		"zhangsan": Person{Age: 20},
	}
	m["zhangsan"].Age = 23
	m["zhangsan"].GrowUp()
}
/*
map 回传的是新的值，修改不会影响到原本的

要改成

1. 修改完后覆盖

p := m["zhangsan"] // tmp
p.Age = 23
p.GrowUp()
m["zhangsan"] = p

2. 直接用指针

m := map[string]*Person{
  "zhangsan": &Person{Age: 20},
}
*/