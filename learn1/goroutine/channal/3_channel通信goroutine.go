package main

import (
	"fmt"
)


//1.什么是协程  用户态的线程调度   由用户创建 go的运行时调度
//2.线程  用户态-内核态的切换
//3.goroutine的 交互channal  两个创新核心 特色
func hello(c chan bool) {
	fmt.Println("hello goroutine")
	c <- true
}

func main() {
	var ex chan bool
	ex = make(chan bool)
	go hello(ex)
	<-ex

}
