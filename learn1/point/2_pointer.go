package main

import "fmt"

func testPoint(){
	var a int = 200
	var b *int = &a

	fmt.Printf("b:%d", *b)

 }
func main() {
	testPoint()
}















