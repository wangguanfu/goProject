package main

import (
	"fmt"
	"runtime"
)

//多核控制
func main() {
	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(1)


}
