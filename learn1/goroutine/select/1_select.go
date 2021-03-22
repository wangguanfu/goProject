package main

import (
	"fmt"
	"time"
)

//多个chan   读取
func sever1(ch chan string) {
	time.Sleep(time.Second * 6)
	ch <- "from sever1"
}

func sever2(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "from sever2"
}

func main() {
	out1 := make(chan string)
	out2 := make(chan string)
	go sever1(out1)
	go sever2(out2)

	//s1 := <-out1
	//fmt.Println("s1",s1)
	//s2 := <-out2
	//fmt.Println("s2",s2)

	select {
	case s1 := <-out1:
		fmt.Println("s1", s1)
	case s2 := <-out2:
		fmt.Println("s2",s2)

	}
}
