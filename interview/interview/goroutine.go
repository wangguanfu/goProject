/*
交替打印数字和字母

	解题思路：
	利用channel控制打印
*/
package main

import (
	"fmt"
	"strings"
	"sync"
)

func main(){
	letter, nums :=make(chan bool), make(chan bool) //利用channel通知
	wait := sync.WaitGroup{}
	go func() {
		i:=1
		for{
			select{
			case <-nums:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter<-true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str:="ABCDEFGHIJKLMNOPQRSTUVWSYZ"
		i:=0
		for{
			select {
			case <-letter:
				if i>=strings.Count(str,"")-1{
					wait.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i++
				nums<-true
				break
			default:
				break
			}
		}
	}(&wait)
	nums<-true
	wait.Wait()
}









