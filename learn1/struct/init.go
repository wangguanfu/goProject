package main

import "fmt"

// go的面向对象 没有构造函数 需要自己实现
type stu struct {
	name string
	age int
	sex string

}

func NewStudents(name string, sex string, age int) *stu{
	user := &stu{
		name: name,
		age: age,
		sex:sex,
	}
	return user
}

func main() {
	s := NewStudents("wang", "1",1)
	fmt.Println(s)
}










