package main

import "fmt"

/*
翻转字符串  不使用额外的数据结构和存储空间
解题思路：
	利用字符指针 使str[len]=str[0]
	以字符串长度1/2为轴 前后赋值
*/

func reverStrings(s string)(string, bool){
	str:=[]rune(s)
	l:= len(str)
	if l>5000{
		return s, false
	}
	for i:=0; i<l/2; i++{
		str[i], str[l-i-1] = str[l-i-1], str[i]

	}
	return string(str),true
}

func main() {
	s := "abcdefg"
	fmt.Println(reverStrings(s))
}






















