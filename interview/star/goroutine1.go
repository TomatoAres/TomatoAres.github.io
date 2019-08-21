package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
