package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func someHander() {
	ctx, cancle := context.WithCancel(context.Background())
	go doStuff(ctx)
	time.Sleep(10*time.Second)
	cancle()
}

func doStuff(ctx context.Context)  {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Println("work")

		}
	}

}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHander()
	logg.Printf("down")
}


/*
1.微服务概述
	1.单体架构
        1.1
		所有的业务功能都在一个应用程序里面
		研发人员共同维护一个代码库
		简单的机构
		1.2
		单体架构的扩展  横向的扩展()

		劣势：
			复杂度高 代码庞大
			开发效率低
			牵一发而动全身

	2.什么是微服务
		微小的服务和应用： 让每个服务专注做自己的事情，
		不同的服务单独开发和部署 服务之间完全隔离

		缺点：
			复杂度高 链路长
			监控和治理

		配套设施
		配套组织架构

2.微服务生态
	硬件层（云 k8s）
		物理服务管理 操作系统 配置 资源隔离和抽象 主机监控和日志
	通信层（tcp  rpc）
		网络传输   RPC  服务发现   服务注册  负载均衡
		消息传递 ：json protobuf thrift

	应用平台层（服务 日志 管理 监控）

    微服务层

3. 全貌
一。微服务框架

二。服务注册
三。服务发现
四。负载均衡
五。rpc
六。自动代码生产器
七。rpc客户端开发



4.CAP



5.raft算法：保证数据一致性  基于复制
	选主
	复制状态机 ：保证日志写入序列一致
	心跳和超时机制
	Raft是一种用来管理日志复制的非对称式一致性算法，相比paxos而言，更容易理解。
    脑裂：网络问题



6.RPC调用
	数据传输
		.thrift
		.protobuf
		.Json

	负载均衡
		.随机
		。轮询
		。一致性哈希
	异常容错
		。健康检查
		。熔断
		。限流

6.服务监控
	。日志收集
	。Metrics打点

233 - 316
*/







