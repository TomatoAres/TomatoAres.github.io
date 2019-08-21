package main

import "fmt"

type duration int

func (d duration) pretty() string {
//func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", d)
}

func question() {
	duration(42).pretty()
}

func improve_1()  {
	d := duration(42)
	d.pretty()
}

func main() {
	s := improve_1
	fmt.Println(s)
}