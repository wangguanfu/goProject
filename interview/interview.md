1. go channel close 

	https://www.zhihu.com/search?type=content&q=go%20channel%20close%20
	
	
2.linux grep 命令查找日志相关文件
tail -500 catalina.out| grep -H -n -B 1 "GenericJDBCException"[/color]	
    https://blog.csdn.net/liupeifeng3514/article/details/79880878	
	
3. b+树结构 and  为什么	
	在计算机中，磁盘存储数据最小单元是扇区，一个扇区的大小是512字节，而文件系统（例如XFS/EXT4）的最小单元是块，一个块的大小是4k，而对于InnoDB存储引擎也有自己的最小储存单元，页（Page），一个页的大小是16K。
	表中的数据都是存储在页中的，所以一个页中能存储多少行数据呢？
	
	InnoDB存储引擎的最小存储单元是页，页可以用于存放数据也可以用于存放键值+指针，在B+树中叶子节点存放数据，非叶子节点存放键值+指针。
    索引组织表通过非叶子节点的二分查找法以及指针确定数据在哪个页中，进而在去数据页中查找到需要的数据；
	
	https://www.zhihu.com/search?type=content&q=b%2B%E6%A0%91%E7%BB%93%E6%9E%84%20and%20%20%E4%B8%BA%E4%BB%80%E4%B9%88
	
	
	由于B+树内部节点不需要存储data pointer，所以B+树单一节点存储的元素更多，使得查询的IO次数更少所有的查询都需要查找到叶节点，查询性能是稳定的，而B树每个节点都可能查询到数据，所以不稳定。所有的叶子节点形成了一个有序链表，能更好地支持范围查询。


3 .slice和array
	go语言的数组是一个值，Go 语言中的数组是值类型
	但他们之间还有着千次万缕的联系 slice 是引用类型、是 array 的引用，相当于动态数组， 这些都是 slice 的特性
	https://www.zhihu.com/search?type=content&q=go%20slice%E5%92%8Carry%E5%8C%BA%E5%88%AB


4. golang中,new和make区别

在golang中,make和new都是分配内存的，但是它们之间还是有些区别的，只有理解了它们之间的不同，才能在合适的场合使用。

简单来说,new只是分配内存，不初始化内存； 而make即分配又初始化内存。所谓的初始化就是给类型赋初值，比如字符为空，整型为0, 逻辑值为false等。
make 仅用来分配及初始化类型为 slice、map、chan 的数据。new 可分配任意类型的数据. new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type. new 分配的空间被清零, make 分配空间后，会进行初始化.


5. IO多路复用 epoll和select区别
	
	https://www.zhihu.com/search?type=content&q=IO%E5%A4%9A%E8%B7%AF%E5%A4%8D%E7%94%A8%20epoll%E5%92%8Cselect%E5%8C%BA%E5%88%AB
	
1、什么是IO多路复用「定义」IO多路复用是一种同步IO模型，
实现一个线程可以监视多个文件句柄；一旦某个文件句柄就绪，就能够通知应用程序进行相应的读写操作；没有文件句柄就绪时会阻塞应用程序，交出cpu。
多路是指网络连接，复用指的是同一个线程

6.进程通信方法
管道
消息队列
共享内存
信号量
信号
套接字

7. tcp保证可靠
02 保证数据安全的方法

https://www.zhihu.com/search?type=content&q=tcp%E4%BF%9D%E8%AF%81%E5%8F%AF%E9%9D%A0

TCP主要提供了检验和、序列号/确认应答、超时重传、最大消息长度、滑动窗口控制等方法实现了可靠性传输。
	

8. 数据库优化

https://www.zhihu.com/search?type=content&q=%E6%95%B0%E6%8D%AE%E5%BA%93%E4%BC%98%E5%8C%96	
	
	
Linux命令——ps、grep以及ps、lsof和netstat的区别


ps（Process status）显示当前进程状态
ps -u root //显示root进程用户信息
ps -ef 可列出所有的进程

区别：
ps ：（process）查看有终端控制的所有进程
ps -ef | grep word 查看包含关键字word的所有外部进程（包括其他用户）
ps -ax 使用 -a 参数。-a 代表 all。同时加上x参数会显示没有控制终端的进程
netstat：用于显示各种网络相关信息，如网络连接，路由表，接口状态
losf：（列出打开文件lists openfiles）能看到pid和用户(有权限控制，只能看到本用户)，可以找到哪个进程占用了这个端口

grep
1、查找文件里符合条件的字符串
2、以递归的方式查找符合条件的文件
grep -r update usr/bin：查找指令目录（usr/bin）及其子目录下，文件中包含字符串“update”的文件，并打印出该字符串所在的行
3、反向查找，通过-v可以查找出不符合条件的行
grep -v test test：查找文件名中包含 test 的文件中不包含test 的行
 
	
	
9. 数据库优化
	redis ACID
	redis  不支持回滚 不能保证持久性  其他都可以
	
	redis最佳实践： 
		内存 性能 高可靠 日常运维 资源规划 监控 安全
		
		1.节省内存：
			控制key的长度
			避免bigkey
			选择合适的数据类型
			把redis当缓存使用 尽可能的合理设置过期时间
			设置合理的淘汰策略
			数据压缩
			
		2.性能
			bigkey
			开启lazy-free
			不使用复杂度过高的命令
			执行o(n)命令时 关注N大小
			使用批量操作
			使用管道
			
		3.高可靠
			按业务线部署实例（资源隔离）
			部署主从集群
			哨兵实例
		
		
		定期清除+懒惰清除

















		
	
	