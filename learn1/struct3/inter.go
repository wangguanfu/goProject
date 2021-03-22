package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	CountInto(&lst, 1, 10)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
		fmt.Println(1111111)
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
// 1.
//评论模块  表设计  -  接口设计
// 评论表  关系表
// 接口  发表评论 发表回复 获取评论列表  获取回复列表  点赞

//2.
//收藏功能  收藏夹 一层目录
// 表设计  收藏夹 id dir_id dir_name  收藏表
// 接口设计 添加收藏夹  list  添加收藏 list


// 3.
//搜索模块
//  etcd   ElasticSearch
//  etcd 配置共享和服务发现  高可用分布式key-value 开发语言Go
// 接口提供restful和http接口 使用简单   基于raft算法的强一致性

// 3.1
// kafka : 1.异步化 提高系统的响应时间和健壮性 2.应用解耦 3.流量削峰
// zookeeper: 1.服务注册和服务发现 注册中心 watch(监听) 2.配置中心 3.分布式锁
// 生产者 - topic - 消费者


//  3.2
//  ElasticSearch
//



























