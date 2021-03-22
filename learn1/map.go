package main

import (
	"fmt"
	"sort"
)

//内建函数 make 用来为 slice，map 或 chan 类型分配内存和初始化一个对象(注意：只能用在这三种类型上)，
//跟 new 类似，第一个参数也是一个类型而不是一个值，跟 new 不同的是，make 返回类型的引用而不是指针

func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["bob"] = 4
	scene["lili"] = 90

	for k, v :=range scene {
		fmt.Println(k, v)
	}
	var sceneList []string
	// 将map数据遍历复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	// 输出
	fmt.Println(sceneList)
}

