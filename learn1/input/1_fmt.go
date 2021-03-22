package main

import "fmt"

func testInput() {
	var a int
	var b string
	var c float32
	fmt.Scanf("%d ,%s, %f", &a, &b, &c)
	fmt.Printf("a , b ,c ", a, b, c)
}

func main() {
	//Scanf Scan ScanLn
	testInput()

}
