package main

import "fmt"

type Address1 struct {
	province string
	city string
}

type S1 struct {
	username string
	age int
	city string
	*Address1
}

func main() {
	user := &S1{
		username: "wang",
		age: 18,
		Address1: &Address1{
			city : "7777",
		},
	}

	fmt.Println(user)
	fmt.Println(user.Address1.city)

}


//二者异同
//二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；而new用于类型的内存分配，并且内存置为零。所以在我们编写程序的时候，就可以根据自己的需要很好的选择了。
//
//make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。






















