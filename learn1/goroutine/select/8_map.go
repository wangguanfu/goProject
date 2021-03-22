package main

//type mapKey struct{
//	key int
//}

//func main() {
//	var m = make(map[mapKey]string)
//	var key = mapKey{10}
//	m[key] = "hello"
//	fmt.Printf("m[key]=%s\n", m[key]) // 修改key的字段的值后再次查询map，无法获取刚才add进去的值    key.key = 100    fmt.Printf("再次查询m[key]=%s\n", m[key])}
//}


//func main() {
//	var m = make(map[string]int)
//	m["a"] = 0
//	fmt.Printf("a=%d; b=%d\n", m["a"], m["b"])
//
//	av, aexisted := m["a"]
//	bv, bexisted := m["b"]
//	fmt.Printf("a=%d, existed: %t; b=%d, existed: %t\n", av, aexisted, bv, bexisted)
//}
//type Counter1 struct {
//	Website      string
//	Start        time.Time
//	PageCounters map[string]int
//}
//
//func main() {
//	var c Counter1
//	c.Website = "baidu.com"
//
//
//	c.PageCounters["/"]++
//}



func main() {
	var m = make(map[int]int,10) // 初始化一个map
	go func() {
		for {
			m[1] = 1 //设置key
		}
	}()

	go func() {
		for {
			_ = m[2] //访问这个map
		}
	}()
	select {}
}

