package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		//sc := bufio.NewScanner(r)
		lines++
		if err != nil {
			break
		}
	}
	if err !=io.EOF{
		return 0, err
	}
	return lines, err
}

func main() {
	var err1 = errors.New("111")
	fmt.Println(err1)
}
