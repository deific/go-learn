package main

import (
	"fmt"
	"time"
)

func ShowFunc()  {
	time.Sleep(time.Second * 3)
	fmt.Println("sleeper finished!")
}

func ShowFunc1(c chan string)  {
	time.Sleep(time.Second * 1)
	fmt.Println("sleeper finished!")
	c <- "sleeper finish1"
	c <- "sleeper finish2"
	// 关闭通道
	close(c)
}
// 方法对通道可读写
func ShowFunc2(c chan string) {
	// 阻塞进程直到收到消息
	for msg := range c {
		fmt.Println("main finished!", msg)
	}
}
// 只读通道
func ShowFunc3(c <- chan string) {
	time.Sleep(time.Second * 1)

}
// 只写通道
func ShowFunc4(c chan <- string) {
}

func DoChan1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping chan1"
}
func DoChan2(c chan  string) {
	time.Sleep(time.Second * 1)
	c <- "ping chan2"
}
func DoChan3(c chan  string) {
	time.Sleep(time.Second * 2)
	c <- "ping chan3"
}

func DoStop(c chan bool) {
	time.Sleep(time.Second * 3)
	//c <- true
	c <- false
}
