package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var done = make(chan struct{}) //channel是多协程安全的，多个goroutinue向channel里面写值，是安全的,channel要初始化
// 很多时候，我并不会多个goroutinue写同一个channel，意味着不同的Goroutinue会
func g1(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(1 * time.Second)
	ch <- struct{}{}
}

func main() {
	g1Channel := make(chan struct{}) //有自己的channel
	g2Channel := make(chan struct{})
	go g1(g1Channel)
	go g2(g2Channel)
	//我要监控多个channel，任何一个channel有值，或者返回都知道
	//1.某一个分支就绪了，就执行该分支
	//2.如果两个都就绪了，先执行哪个？这个是随机的，随机的原因，目的是防止饥饿
	//3.应用场景:
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done...")
		case <-g2Channel:
			fmt.Println("g2 done...")
		case <-timer.C:
			fmt.Println("timeout...")
			//time.Sleep(10 * time.Millisecond)
			return
		}
	}

}
