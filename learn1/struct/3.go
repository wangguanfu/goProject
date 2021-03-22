package main

import "fmt"

// struct  是一块连续的内存布局
type test struct {
	a int32
	b int32
	c int32

}

func main() {
	var t test
	fmt.Printf("%p\n", &t.a)
	fmt.Printf("%p\n", &t.b)
	fmt.Printf("%p\n", &t.c)

}



