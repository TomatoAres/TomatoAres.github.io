package main

type Test struct {}

func main() {

	//var t Test
	//println(t == nil)//cannot convert nil to type Test

	var t2 *Test
	println(t2 == nil)

	println(func()bool{ var t2 *Test;return (t2== nil)}())
}
