package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}














