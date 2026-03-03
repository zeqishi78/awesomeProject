package main

import (
	"fmt"
	"time"
)

/*
单向channel
*/
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num = %d\r\n", num)
	}
}

func main() {
	//默认情况下，channel是双向的，可以发数据，也可以接收数据，但是我们经常一个channel，作为参数进行传递，这时候要考虑单向的功能了
	//var ch1 chan int       //这是双向channel
	//var ch2 chan<- float64 //单向channel，只能写入float64的数据
	//var ch3 <-chan int     //单向的，只能读取数据
	//
	//c := make(chan int, 3)
	//var send chan<- int = c //send-only
	//var read <-chan int = c //reev-only
	//
	//send <- 1 //存值的时候，把右边的值放在左边

	//d1 := (chan int)(send) //不能把单向的channel，转换为双向的channel，Cannot convert an expression of the type 'chan<- int' to the type 'chan int'
	ch := make(chan int)
	go producer(ch)

	go consumer(ch)
	time.Sleep(10 * time.Second)

}
