package main

import "fmt"

func countBinarySubstrings(s string) int {
	counts := []int{}
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		counts = append(counts, count)
		fmt.Println(counts)
	}
	ans := 0
	for i := 1; i < len(counts); i++ {
		fmt.Println(counts[i], counts[i-1])
		ans += min(counts[i], counts[i-1])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s := "00110011"
	countBinarySubstrings(s)
}