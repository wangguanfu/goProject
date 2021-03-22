package main

import (
	"fmt"
	"time"
)

//1.什么是协程  用户态的线程调度   由用户创建 go的运行时调度
//2.线程  用户态-内核态的切换
//3.goroutine的 交互channal  两个创新核心 特色
func hello(i int){
	fmt.Println("hello goroutine" ,i)
}

func main() {
	//go hello()
	//fmt.Println(" main")
	//time.Sleep(1) //创建消耗时间  需要主线程等待
	for i:=0; i<10; i++ {
		go hello(i)
	}
	time.Sleep(time.Second)
}





















