package main

import "fmt"

type People struct{}

func (p *People)ShowA()  {
	fmt.Println("Show A")
	p.ShowB()
}

func (p *People)ShowB()  {
	fmt.Println("Show B")
}

type Teacher struct {
	People
}
func (t *Teacher)showB()  {
	fmt.Println("Teacher show B")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

// output:
// Show A
//Show B