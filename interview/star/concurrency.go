package main

import (
	"fmt"
	"sync"
)

/*
有以下代码
```
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
```
输出是hello还是welcome？为什么？

再进一步，有以下代码
```
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	wg.Wait()


```
输出结果是什么？为什么？*/

func t1()  {
	fmt.Println("==========t1:")
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func t2()  {
	fmt.Println("==========t2:")
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	wg.Wait()
}

func main() {
	t1()
	fmt.Println()
	t2()
}
/*
第一个结果：fmt.Println(salutation)语句在协程运行完之后才会运行，此时salutation已经变为welcome了
第二个结果应该是不确定。
for语句中的salution共享同一个变量地址，由于不确定在协程中打印的时候该变量存的是哪个值，所以结果应该是不确定吧。
*/