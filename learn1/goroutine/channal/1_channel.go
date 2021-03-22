package main

import "fmt"

func main() {
	var c chan int
	fmt.Printf("%v\n", c)
	c = make(chan int, 1) //初始化  分匹配地址
	fmt.Printf("%v\n", c)

	c <- 100
	data := <- c
	fmt.Println(data)

}
