package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Println(str)
}



