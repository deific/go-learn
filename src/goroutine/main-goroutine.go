package main

import (
	"fmt"
	"time"
)

func main() {
	sleepControl()
	chanControl()
	bufferChanControl()
	selectChanControl()
	forAndSelectChanControl()
	s := "sss"
	fmt.Printf("string is %v\n", s)
}


// 通过定时器控制goroutine的执行
func sleepControl() {
	// 并行执行
	go ShowFunc()

	fmt.Println("main finished!")
	time.Sleep(time.Second * 4)
}

// 通过通道控制
func chanControl() {
	c := make(chan string)
	// 并行执行
	go ShowFunc1(c)
	// 阻塞进程直到收到消息
	for msg := range c {
		fmt.Println("main finished!", msg)
	}
}

// 缓冲通道控制
func bufferChanControl() {
	// 创建长度为2的缓冲通道，只容许最多2条消息
	c := make(chan string, 2)
	c <- "hello 1"
	c <- "hello 2"
	// 关闭通道
	close(c)
	// 并行执行
	go ShowFunc2(c)

	time.Sleep(time.Second * 4)
}

// 缓冲通道控制
func selectChanControl() {
	// 创建长度为2的缓冲通道，只容许最多2条消息
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go DoChan1(c1)
	go DoChan2(c2)
	go DoChan3(c3)

	// select 只执行一次，且只执行最先返回的通道消息
	select {
	case msg := <-c1:
		fmt.Println("do :", msg)
	case msg := <-c2:
		fmt.Println("do :", msg)
	case msg := <-c3:
		fmt.Println("do :", msg)
	case <-time.After(5000 * time.Millisecond): // 超时机制
		fmt.Println("time is out")
	}

}

// 缓冲通道控制
func forAndSelectChanControl() {
	// 创建长度为2的缓冲通道，只容许最多2条消息
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	stop := make(chan bool)

	go DoChan1(c1)
	go DoChan2(c2)
	go DoChan3(c3)
	go DoStop(stop)

	// 阻塞无限循环控制select执行每个通道的返回
	for {
		// select 只执行一次，且只执行最先返回的通道消息
		select {
		case msg := <-c1:
			fmt.Println("do :", msg)
		case msg := <-c2:
			fmt.Println("do :", msg)
		case msg := <-c3:
			fmt.Println("do :", msg)
		case isStop := <-stop: // 通过专门的一个通道控制for循环的退出
			if isStop {
				fmt.Println("stop now")
				return
			}
		}
	}

}
