package main

import "fmt"

//   定义与介绍
// 1.只定义 不实现   2.具体对象 实现细节

type Anmimal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {

}

func (d Dog) Talk(){
	fmt.Println("wangwang")
}
func (d Dog) Eat() {
	fmt.Println("eat eat")
}
func (d Dog) Name() string {
	fmt.Println("wangcai")
	return "dog"
}

func main() {
	var a Anmimal  //  初始化接口类型
	var d Dog   //  实例化 类

	var d1 *Dog = &Dog{} //值类实现的接口 指针也可以调用 反之不可以

	a = d  //  初始化类的接口

	a.Talk()
	a.Eat()
	a.Name()
	d1.Eat()


}










