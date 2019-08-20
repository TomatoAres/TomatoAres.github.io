package main

import "fmt"

func main() {
	var s_chan = make(chan string)
	var int_chan = make(chan int,1)

	s_chan <- "hello"
	int_chan<- 1
	select {
	case <-s_chan:
		fmt.Println("nihao")
	case <-int_chan:
		panic("invalid")
	}

}
