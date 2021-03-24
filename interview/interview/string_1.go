package main

import "strings"

/*
	判断字符串中字符是否相同
		要求1 不使用额外的存储
		要求2 字符为ACS11字符 字符长度小于3000
*/

func isUnqinueSting(s string)bool{
	if strings.Count(s, "")>3000{
		return false
	}
	for _, v := range s{
		if v>127{
			return false
		}
		if strings.Count(s, string(v))>1{
			return false
		}
	}
	return true
}

func isUnqinueSting2(s string)bool{
	if strings.Count(s, "")>3000{
		return false
	}
	for k, v := range s {
		if v>127{
			return false
		}
		if strings.Index(s, string(v))!=k{
			return false
		}
	}
	return true
}


func main() {




}




















