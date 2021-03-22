package main

import "fmt"

func main() {
	type Command struct {
		Name    string    // 指令名称
		Var     *int      // 指令绑定的变量
		Comment string    // 指令的注释
	}

	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"

	fmt.Println(*cmd.Var)


// 使用键值对填充结构体的例子

	type People struct {
		name  string
		child *People
	}
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	fmt.Println(relation.name, relation.child.name, relation.child.child.name)



	type Address struct {
		Province    string
		City        string
		ZipCode     int
		PhoneNumber string
	}
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)


	fmt.Println(NewBlackCat("blue"))
}


type Cat struct {
	Color string
	Name  string
}
type BlackCat struct {
	Cat  // 嵌入Cat, 类似于派生
}
// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

