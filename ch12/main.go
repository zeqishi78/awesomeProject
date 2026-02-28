package main

import (
	"fmt"
	"time"
)

/*
内存占用小:2k、切换快，go语言的协程，go语言诞生之后就只有协程可用-goroutine，非常方便
*/
func asyncPrint() {
	for {
		time.Sleep(time.Second)
		fmt.Println("bobby123")
	}

}

func main() {
	//主死随从
	go asyncPrint()                //普通函数前面加一个go表示异步执行
	fmt.Println("main goroutinue") //异步打印，这个先执行，然后再等3秒，在执行异步代码
	time.Sleep(10 * time.Second)
}
