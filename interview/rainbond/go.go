package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 考点1
func init() {
	// 考点2 没有这行会发生什么
	// go1.5 以上，默认会为每个可用的物理处理器分配一个逻辑处理器
	runtime.GOMAXPROCS(1)
}

func main() {
	// 考点3
	fmt.Println("concurrence version:",runtime.Version())
	var wg sync.WaitGroup
	count := 10
	wg.Add(count * 2)


	//考点4
	for i:= 0;i < count;i++{
		go func() {
			defer wg.Done()
			fmt.Println("concurrence a:",i)
		}()
	}

	//考点5
	for i:= 0;i <count;i++{
		go func(i int) {
			defer wg.Done()
			fmt.Println("concurrence b:",i)
		}(i)
	}
	wg.Wait()

	//time.Sleep(time.Millisecond)
}
