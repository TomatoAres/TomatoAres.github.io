package main

import (
	"log"
	"os"
	"time"
)

/*func ProcessChannelMessages(ctx context.Context, in <-chan string) {
	for {
		start := time.Now()
		select {
		case s, ok := <-in:
			if !ok {
				return
			}
			// handle `s`
		case <-time.After(5 * time.Minute):
			idleCounter.Inc()
		case <-ctx.Done():
			return
		}
	}
}*/

/*1.start没使用……应该会异常。
2.因为是个for无限循环，所以time.after每次都会重新计算，所以浪费了资源，而且这条语句看起来永远执行不到

关于改进:  个人看法
1. 将开始时间提到for之外，并且在for循环内定义时间，处理两个时间的差值。
但由于select，只接受chan的传值，目前想到只能放在default中，感觉并不太友好。
开始时间*/

func improve_1(in <-chan string)  {
	start := time.Now()
	// 最大延迟
	var maxDuration float64 = 1
	for {
		// 定义当前时间
		now := time.Now()
		select {
		case s, ok := <-in:
			if !ok {
				return
			}
			// handle s
			log.Println(s)
		default :
			if now.Sub(start).Minutes() > maxDuration {
				// do something! bob!
				log.Println("overtime")
				os.Exit(1)
			}
		}
	}
}


/*2.个人改进2:
利用time.tick定时器，但也需要放在for之外

补充：tick 记得 stop
*/
func improve_2(in <- chan string)  {
	tick := time.Tick(5 * time.Second)
	for {
		// 定义当前时间
		select {
		case s, ok := <-in:
			if !ok {
				return
			}
			// handle s
			log.Println(s)
		case <-tick:
			// It's overtime!! please do something!!
			log.Println("overtime")
			os.Exit(1)
		}
	}
}

func main() {
	s := make(chan string,1)
	s<- "hello"
	improve_1(s)

}