package main

import "fmt"

// 面向对象的 组成 成员和方法 （在函数里   定义接收者）

type Peaple struct {
	name string
	country string
}

func(p Peaple) Print(){
	fmt.Println(p.country, p.name)
}

func main() {
	var p1 Peaple = Peaple{
		name: "1",
		country: "china",
	}
	p1.Print()
}

//函数 不属于 任何类型   方法属于类型的（没有方法重载）


























