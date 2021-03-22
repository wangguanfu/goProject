package main

import "fmt"

type User struct {
	name string
	sex string
	age int
}

func main()  {
	var user User
	user.age =18
	user.name = "wang"
	user.sex = "女"

	fmt.Println(user)


	user2 := User{
		name: "wang",
		age:18,
		sex:"男",
	}
	fmt.Println(user2)
}












