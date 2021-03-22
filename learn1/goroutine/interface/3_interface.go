package main

import "fmt"

//接口类型的变量 --- 可以存储所有的实现的变量的实例
//及 空接口  可以是所有类型

func describe(i interface{}) {
	fmt.Printf("type = %T, value = %v\n", i, i)
}

func main() {
	var a interface{}

	var b = 100
	a = b
	fmt.Printf("%T, %v\n", a, a)

	var c = "阿珍爱上了阿强"
	a = c
	fmt.Printf("%T, %v\n", a, a)

	var d = make(map[string]int, 100)
	d["abc"] = 1000
	d["ek"] = 30

	a = d
	fmt.Printf("%T, %v\n", a, a)

	for k, v := range d {
		fmt.Println(k, v)
	}


	s  := "hello world"
	describe(s)

	i:=55
	describe(i)

	strt := struct {
		name string
	}{
		name: "hello",

	}
	describe(strt)

}



