package main

import "fmt"

func main() {


	//变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址。

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	//取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	//
	//变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	//对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
	//指针变量的值是指针地址。
	//对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。


	str := new(string)
	*str = "go 语言教程"
	fmt.Println(*str)


}
