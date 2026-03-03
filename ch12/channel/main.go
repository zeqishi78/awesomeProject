package main

import "fmt"

/*
go语言中，不要通过共享内存来通信，而要通过通信来实现内存共享
php python java、多线程编程的时候，两个Goroutinue之间的通信最常见的方式是一个全局变量，
其他语言中会提供消息队列的机制，python中会提供queue，java也会提供queue，就是消费者和生产者的关系，
生产者把消息往消息队列中放，消费者从消息队列中消费消息。

为了实现这种方式，go提供了channel提供了语法糖，让使用channel的方式更加简单

无缓冲channel，适用于通知，比如说B的goroutinue，等待A的goroutinue的通知，这个通知是事件有没有发生的通知，B要第一时间知道A是否完成，
这种就适用于无缓冲channel，A发出之后，B立马就能收到，很快的
有缓冲channel，适用于消费者和生产者之间的通信，

go中channel的应用场景：

	1.消息传递、消息过滤
	2.信号广播
	3.事件订阅和广播
	4.任务分发
	5.结果汇总
	6.并发控制
	7.同步和异步
	...
*/
func main() {
	/*
		channel底层是一个环形数组来实现的，

	*/
	var msg chan string        //定义一个channel，里面放的数据类型是string类型
	msg = make(chan string, 1) //初始化channel，channel的初始化值，如果为0的话，就是无缓冲的channel
	//mas = make(chan string, 0) //初始化值为0，无缓冲的channel
	msg <- "bobby22" //放一个字符串到channel中
	data := <-msg    //取值
	fmt.Println(data)

	//无缓冲数据处理

}
