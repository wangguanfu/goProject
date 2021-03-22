package main

import (
	"fmt"
	"sync"
	"time"
)


//1.等待一组goroutine结束

func process1(i int, wg *sync.WaitGroup) {
	fmt.Println("start goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i:=0; i<no; i++ {
		wg.Add(1)
		go process1(i, &wg)
	}
	wg.Wait()
	fmt.Println("exit")
}




































