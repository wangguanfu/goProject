package main

import (
	"fmt"
	"os"
)

func main() {
	var buf [16]byte
	os.Stdin.Read(buf[:])
	fmt.Println(string(buf[:]))
}
