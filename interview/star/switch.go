package main

import "fmt"

func False() bool {
	return false
}

/*Switch without a condition is the same as switch true.*/
func switch_1() {
	switch False() {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}
}

func switch_2()  {
	switch False()
	{
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}
}

func main() {
	fmt.Println("switch1")
	switch_1()

	fmt.Println("switch_2")
	switch_2()
}