package main

import "fmt"

type User1 struct {
	name string
	sex string
	age int
}

func main()  {
	var user *User1
	fmt.Printf("user =%v\n", user)

	var user01 = &User1{}
	user01.name = "xinxin"
	fmt.Printf("user =%v\n", user01)


	user02 := &User1{
		name:"xin",
		age:18,
	}
	fmt.Println(user02)

	var user03 = new(User1)
	fmt.Println(user03)
}







