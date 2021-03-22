package main

import "fmt"

func main(){
	var n [10]int
	var i,j int
	for i=0; i<10;i++{
		n[i] = i+100
	}
	for j=0;j<10;j++{
		fmt.Printf("E[%d] = %d\n", j, n[j])
	}

	var a = [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /* 第三行索引为 2 */
	}
	fmt.Println(a)
}




