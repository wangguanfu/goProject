package main

import "fmt"

type Address struct {
	province string
	city string
}

type S struct {
	username string
	age int
	address Address
	city string
}

func main() {
	user := &S{
		username: "wang",
		age: 18,
		city : "123213123",
		address: Address{
			province: "hei",
			city: "a",
		},
	}

	fmt.Println(user.city)
	fmt.Println(user.address.city)

}



