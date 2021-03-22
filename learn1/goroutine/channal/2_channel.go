package main

import (
	"fmt"
	"time"
)

func producer(c chan int) {
	c <- 1000
}

func consumer(c chan int)  {
	data := <- c
	fmt.Println(data)
}

func main() {
	var c chan int
	c = make(chan int)
	go producer(c)
	go consumer(c)

	time.Sleep(time.Second)
}





















