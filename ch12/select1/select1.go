package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var lock sync.Mutex

func g1() {
	time.Sleep(time.Second)
	lock.Lock()
	defer lock.Unlock()
	done = true
}

func g2() {
	time.Sleep(2 * time.Second)
	lock.Lock()
	defer lock.Unlock()
	done = true
}

func main() {
	//select类似switch case，但是select的功能和我们操作linux里面提供的io的select 、poll、epoll用法和功能差不多，
	//select主要作用于多个channel，select语句执行的时候，会选择目前已经就绪的channel
	//现在有个需求，我现在有2个Goroutinue都在执行，但是我在主的goroutinue中，当某个执行完成之后，这个时候我会立马知道，
	go g1()
	go g2()

	for {
		if done {
			fmt.Println("done")
			time.Sleep(10 * time.Millisecond)
			return
		}
	}

}
