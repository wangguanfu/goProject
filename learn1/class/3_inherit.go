package main

import "fmt"

//继承
type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk() {
	fmt.Println("talk")
}

type Dog struct {
	Feet string
	Animal
}

func (d *Dog) Fat() {
	fmt.Println("dog is eat")
}

func main() {
	var d *Dog = &Dog{
		Feet: "furt",
		Animal: Animal{
			Name: "1",
			Sex:  "2",
		},
	}
	fmt.Println(d.Sex)
	d.Fat()
	d.Talk()
}
