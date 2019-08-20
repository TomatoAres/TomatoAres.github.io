package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const N = 26

func main() {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)
	var wg sync.WaitGroup
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c", 'a'+i)
		}(i)

		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c", 'A'+i)
		}(i)
	}
	time.Sleep(1e9)
	wg.Wait()
}
