package main

import "fmt"

/*
	滑动窗口
		滑动窗口是双指针的一种 利用双指针在数组单一方向滑动 形成子区间
		对子区间进行剪纸 最终得到满足条件的

*/

//无重复最长字串
func lengthOfLongestSubString(s string) int {
	lens := len(s)
	windows := make(map[byte]int, lens)
	left, right, res := 0, 0, 0
	for right < lens {
		b := s[right]
		windows[b]++
		right++
		for windows[b] > 1 {
			c := s[left]
			windows[c]--
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

func main() {
	a := "abcabcbb"
	fmt.Println(lengthOfLongestSubString(a))
}