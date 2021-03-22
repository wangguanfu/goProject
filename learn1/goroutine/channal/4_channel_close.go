package main

import "fmt"

func pro(c chan int ){
	for i:=0; i<10;i++ {
		c <- i
	}
	close(c)
}

func main() {
	ch := make(chan int)
	go pro(ch)
	for {
		v, ok := <-ch
		if ok == false{
			break
		}
		fmt.Println(v, ok)
	}
}
























