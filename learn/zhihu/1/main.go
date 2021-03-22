package main

import "fmt"

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func main() {
	// 1.

	//s := make([]int, util)
	//fmt.Println(s)
	//
	//s = append(s, 1, 2, 3)
	//fmt.Println(s)


	// 2.

	s := make([]int,0)
	fmt.Println(s)

	s = append(s,1,2,3,4)
	fmt.Println(s)

}