package main

import "fmt"

func pro1(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	ch := make(chan int)
	go pro1(ch)
	for v := range ch {
		fmt.Println("ver", v)
	}
}
