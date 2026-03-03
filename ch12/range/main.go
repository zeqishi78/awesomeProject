package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int
	msg = make(chan int, 0)
	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all done")
		//data := <-msg
		//fmt.Println(data)
		//data = <-msg
		//fmt.Println(data)
		//data = <-msg
		//fmt.Println(data)
	}(msg)

	msg <- 1
	msg <- 2
	msg <- 5
	close(msg) //其他的编程语言有很大区别
	//msg <- 3   //放值到channel中，已经关闭的channel不能再放值了
	d := <-msg
	fmt.Printf("d = %d\r\n", d)
	time.Sleep(time.Second * 10)
}
