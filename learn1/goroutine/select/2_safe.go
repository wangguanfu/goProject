package main

import (
	"fmt"
	"time"
)

//线程安全  同时操作一个资源

var x int

func add() {
	for i := 0; i < 10000; i++ {
		x = x + 1
	}

}

func main() {
	go add()
	go add()
	time.Sleep(time.Second)
	fmt.Println(x)
}
