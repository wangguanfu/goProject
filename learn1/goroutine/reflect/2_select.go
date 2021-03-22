package main

import (
"fmt"
"reflect"
)

func main() {
	c := make(chan int, 1)
	vc := reflect.ValueOf(c)
	succeeded := vc.TrySend(reflect.ValueOf(123))
	fmt.Println(succeeded, vc.Len(), vc.Cap())

	vSend, vZero := reflect.ValueOf(789), reflect.Value{}
	branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault, Chan: vZero, Send: vZero},
		{Dir: reflect.SelectRecv, Chan: vc, Send: vZero},
		{Dir: reflect.SelectSend, Chan: vc, Send: vSend},
	}

	// use of Select() method
	selIndex, vRecv, sentBeforeClosed := reflect.Select(branches)
	fmt.Println(selIndex)
	fmt.Println(sentBeforeClosed)
	fmt.Println(vRecv.Int())
	vc.Close()
	// Remove the send case branch this time,
	// for it may cause panic.
	//selIndex, _, sentBeforeClosed = reflect.Select(branches[:2])
	//fmt.Println(selIndex, sentBeforeClosed)
}
