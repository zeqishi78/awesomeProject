package main

import (
	"fmt"
	"time"
)

/*
	使用两个goroutinue交替打印序列，一个goroutinue打印数字，另外一个goroutinue打印字母，最终效果如下：
		12AB34CD56EF78GH910TJ1112KL1314MN15160P17180R1928ST2122UV2324WX2526YZ2728
	这段代码是Go语言中，利用Channel(通道)和Goroutinue(协程)实现并发控制，核心目的是让不同的协程按照数字-字母-数字-字母的顺序交替执行，从而生成特定的字符串序列
	核心机制：Channel作为信号量
*/
//定义两个全局的无缓冲Channel
var number, letter = make(chan bool), make(chan bool)

/*
流程:等待-->打印-->通知对方-->循环
*/
func printNumber() {
	i := 1
	for { //关键点1:等待信号
		<-number
		//只有当main函数获取其他协程向number通道发送了true，这里才会继续执行
		//怎么做到先打印数字，此处应该怎么做到交叉打印，应该等待另一个goroutinue来通知我来打印了
		fmt.Printf("%d%d", i, i+1) //打印两个连续数字，12、34
		i += 2                     //数字递增
		letter <- true             //关键点2：发送信号
		//向letter通道发送true，唤醒printLetter协程，自己在此处阻塞直到对方接收
	}
}

/*
 */
func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-letter //[关键点1]等待信号
		//只有当printNumber向letter通道发送了true，这里才会继续发送
		if i >= len(str) {
			return //如果字母用完了，直接退出协程
		}
		fmt.Print(str[i : i+2]) //连续打印两个字母，如AB、CD
		i += 2                  //索引递增
		//向Number通道发送true，唤醒printNumber协程
		number <- true
	}
}

/*
启动顺序:

	1.两个字协程启动，但是都因为等待通道数据而挂起
	2.main函数向number通道写入true
	3.printNumber被唤醒，打印12，然后向letter通道写入true并挂起
	4.printLetter被唤醒，打印AB，然后向number写入true并挂起。
	5.如此循环往复，形成12-->AB-->34-->CD..的交替输出
*/
func main() {
	go printNumber() //启动数字协程，此时它卡在<-number处等待
	go printLetter() //启动字母协程，此时它卡在<-letter处等待
	number <- true   //启动
	//主协程向number通道发送第一个信号，这会让printNumber中的<-number解除阻塞，开始第一次打印
	time.Sleep(time.Second * 100)
}
