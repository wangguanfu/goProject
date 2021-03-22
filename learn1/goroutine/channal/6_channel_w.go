package main

import (
	"fmt"
	"time"
)

//1.等待一组goroutine结束

func process(i int, c chan bool) {
	fmt.Println("start goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d\n", i)
	c <-true
}

func main() {
	no := 3
	ex := make(chan bool , no)
	for i:=0; i<no; i++ {
		go process(i, ex)
	}
	for i:=0; i<no; i++ {
		<- ex
	}
	fmt.Println("exit")
}

