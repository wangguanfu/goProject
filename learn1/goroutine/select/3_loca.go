package main

import (
	"fmt"
	"sync"
)

//线程安全  同时操作一个资源

var y int
var mutex sync.Mutex
var wg sync.WaitGroup

func add1() {
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		y = y + 1
		mutex.Unlock()
	}
	wg.Done()

}

func main() {
	wg.Add(2)
	go add1()
	go add1()
	wg.Wait()
	fmt.Println(y)
}
