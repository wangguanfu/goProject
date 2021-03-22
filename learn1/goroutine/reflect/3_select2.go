package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var sendCh = make(chan int)

	var increaseInt = func(c chan int) {
		for i := 0; i < 8; i++ {
			c <- i
		}
		close(c)
	}

	go increaseInt(sendCh)

	var selectCase = make([]reflect.SelectCase, 1)
	selectCase[0].Dir = reflect.SelectRecv
	selectCase[0].Chan = reflect.ValueOf(sendCh)

	counter := 0
	for counter < 1 {

		// use of Select() method
		chosen, recv, recvOk := reflect.Select(selectCase)
		if recvOk {
			errors.New("111")
			//errors.Unwrap("chosen, recv, recvOk ",chosen, recv, recvOk )
			fmt.Println(chosen, recv.Int(), recvOk)
			errors.Is()
		} else {
			counter++
		}
	}
}
